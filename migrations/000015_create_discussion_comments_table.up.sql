CREATE TABLE IF NOT EXISTS discussion_comments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    post_id UUID NOT NULL REFERENCES discussion_posts(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    parent_comment_id UUID REFERENCES discussion_comments(id) ON DELETE CASCADE,
    body TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_discussion_comments_post_id ON discussion_comments(post_id);
CREATE INDEX idx_discussion_comments_parent_id ON discussion_comments(parent_comment_id);
CREATE INDEX idx_discussion_comments_user_id ON discussion_comments(user_id);
