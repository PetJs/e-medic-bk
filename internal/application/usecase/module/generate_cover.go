// Package module contains module management use cases.
package module

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"mime"
	"time"

	"emedic-bk/internal/domain/repository"
	"emedic-bk/internal/domain/service"
)

// ModuleCoverGenerator generates an AI cover image for a module in the
// background from its title and description, and stores the result in R2.
type ModuleCoverGenerator struct {
	moduleRepo repository.ModuleRepository
	storage    service.StorageService
	imageGen   service.ImageGenerator
}

// NewModuleCoverGenerator creates a new ModuleCoverGenerator. imageGen may be
// nil (e.g. no API key configured yet), in which case Trigger is a no-op.
func NewModuleCoverGenerator(
	moduleRepo repository.ModuleRepository,
	storage service.StorageService,
	imageGen service.ImageGenerator,
) *ModuleCoverGenerator {
	return &ModuleCoverGenerator{moduleRepo: moduleRepo, storage: storage, imageGen: imageGen}
}

// Trigger kicks off cover generation in the background; callers do not wait for it.
func (g *ModuleCoverGenerator) Trigger(moduleID, title, description string) {
	if g == nil || g.imageGen == nil {
		return
	}
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
		defer cancel()

		data, mimeType, err := g.imageGen.GenerateImage(ctx, buildCoverPrompt(title, description))
		if err != nil {
			slog.Error("module cover generation failed", "module_id", moduleID, "error", err)
			_ = g.moduleRepo.UpdateCoverImage(ctx, moduleID, "", "failed")
			return
		}

		key := fmt.Sprintf("modules/%s/cover%s", moduleID, extForMimeType(mimeType))
		if err := g.storage.Upload(ctx, key, bytes.NewReader(data), mimeType, int64(len(data))); err != nil {
			slog.Error("module cover upload failed", "module_id", moduleID, "error", err)
			_ = g.moduleRepo.UpdateCoverImage(ctx, moduleID, "", "failed")
			return
		}

		if err := g.moduleRepo.UpdateCoverImage(ctx, moduleID, key, "ready"); err != nil {
			slog.Error("module cover status update failed", "module_id", moduleID, "error", err)
		}
	}()
}

func buildCoverPrompt(title, description string) string {
	return fmt.Sprintf(
		"A minimalist, flat-design vector illustration representing the medical education "+
			"topic %q. %s. Clean geometric shapes, calm teal (#0D9488) and warm amber (#D97706) "+
			"color palette on a soft white background, professional educational style suitable "+
			"for a university medical students' e-learning platform. No text, letters, or "+
			"numbers in the image. Wide 16:9 composition.",
		title, description,
	)
}

func extForMimeType(mimeType string) string {
	exts, err := mime.ExtensionsByType(mimeType)
	if err != nil || len(exts) == 0 {
		return ".png"
	}
	return exts[0]
}
