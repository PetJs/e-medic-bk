-- Enforce at most one active subscription per user.
-- First cancel duplicates (keep the one with the latest period end),
-- then add a partial unique index so concurrent activations cannot race.
UPDATE subscriptions s
SET status = 'canceled', canceled_at = NOW(), updated_at = NOW()
WHERE s.status = 'active'
  AND EXISTS (
    SELECT 1 FROM subscriptions k
    WHERE k.user_id = s.user_id
      AND k.status = 'active'
      AND (k.current_period_end > s.current_period_end
           OR (k.current_period_end = s.current_period_end AND k.id > s.id))
  );

CREATE UNIQUE INDEX idx_subscriptions_one_active ON subscriptions(user_id) WHERE status = 'active';
