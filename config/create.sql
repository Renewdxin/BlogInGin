-- noinspection SqlCurrentSchemaInspectionForFile

CREATE DATABASE blog_server;


-- created_on INT DEFAULT '0',
--     created_by VARCHAR(100) DEFAULT '',
--     modified_on INT DEFAULT '0',
--     modified_by VARCHAR(255) DEFAULT '',
--     deleted_on INT DEFAULT '0',
--     is_del SMALLINT DEFAULT '0',

CREATE TABLE IF NOT EXISTS blog_tag (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) DEFAULT '',
    created_on INT DEFAULT '0',
    created_by VARCHAR(100) DEFAULT '',
    modified_on INT DEFAULT '0',
    modified_by VARCHAR(255) DEFAULT '',
    deleted_on INT DEFAULT '0',
    is_del SMALLINT DEFAULT '0',
    state SMALLINT DEFAULT '1'
);

COMMENT ON COLUMN blog_tag.name IS '标签名称';
COMMENT ON COLUMN blog_tag.created_on IS '创建时间';
COMMENT ON COLUMN blog_tag.created_by IS '创建人';
COMMENT ON COLUMN blog_tag.modified_on IS '修改时间';
COMMENT ON COLUMN blog_tag.modified_by IS '修改人';
COMMENT ON COLUMN blog_tag.deleted_on IS '删除时间';
COMMENT ON COLUMN blog_tag.is_del IS '是否删除 0 为未删除 1 为删除';
COMMENT ON COLUMN blog_tag.state IS '状态 0 为禁用 1 为启用';

CREATE TABLE IF NOT EXISTS blog_article (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) DEFAULT '',
    "desc" VARCHAR(255) DEFAULT '',
    cover_image_url VARCHAR(255) DEFAULT '',
    content TEXT DEFAULT '',
    created_on INT DEFAULT '0',
    created_by VARCHAR(100) DEFAULT '',
    modified_on INT DEFAULT '0',
    modified_by VARCHAR(255) DEFAULT '',
    deleted_on INT DEFAULT '0',
    is_del SMALLINT DEFAULT '0',
    state SMALLINT DEFAULT '1'
);

COMMENT ON COLUMN blog_article.title IS '文章标题';
COMMENT ON COLUMN blog_article."desc" IS '文章描述';
COMMENT ON COLUMN blog_article.cover_image_url IS '文章封面图片';
COMMENT ON COLUMN blog_article.content IS '文章内容';
COMMENT ON COLUMN blog_article.created_on IS '创建时间';
COMMENT ON COLUMN blog_article.created_by IS '创建人';
COMMENT ON COLUMN blog_article.modified_on IS '修改时间';
COMMENT ON COLUMN blog_article.modified_by IS '修改人';
COMMENT ON COLUMN blog_article.deleted_on IS '删除时间';
COMMENT ON COLUMN blog_article.is_del IS '是否删除 0 为未删除 1 为删除';
COMMENT ON COLUMN blog_article.state IS '状态 0 为禁用 1 为启用';

-- noinspection SqlCurrentSchemaInspection
CREATE TABLE IF NOT EXISTS blog_article_tag (
    id SERIAL PRIMARY KEY,
    article_id INT DEFAULT '0',
    tag_id INT DEFAULT '0',
    created_on INT DEFAULT '0',
    created_by VARCHAR(100) DEFAULT '',
    modified_on INT DEFAULT '0',
    modified_by VARCHAR(255) DEFAULT '',
    deleted_on INT DEFAULT '0',
    is_del SMALLINT DEFAULT '0'
);

COMMENT ON COLUMN blog_article_tag.article_id IS '文章id';
COMMENT ON COLUMN blog_article_tag.tag_id IS '标签id';
COMMENT ON COLUMN blog_article_tag.created_on IS '创建时间';
COMMENT ON COLUMN blog_article_tag.created_by IS '创建人';
COMMENT ON COLUMN blog_article_tag.modified_on IS '修改时间';
COMMENT ON COLUMN blog_article_tag.modified_by IS '修改人';
COMMENT ON COLUMN blog_article_tag.deleted_on IS '删除时间';
COMMENT ON COLUMN blog_article_tag.is_del IS '是否删除 0 为未删除 1 为删除';


COMMENT ON TABLE blog_article IS '文章管理';
COMMENT ON TABLE blog_article_tag IS '文章标签关联表';


CREATE TABLE IF NOT EXISTS blog_auth (
    id SERIAL NOT NULL PRIMARY KEY ,
    app_key VARCHAR(20) DEFAULT '' ,
    app_secret VARCHAR(50) DEFAULT '',
    created_on INT DEFAULT '0',
    created_by VARCHAR(100) DEFAULT '',
    modified_on INT DEFAULT '0',
    modified_by VARCHAR(255) DEFAULT '',
    deleted_on INT DEFAULT '0',
    is_del SMALLINT DEFAULT '0'
);