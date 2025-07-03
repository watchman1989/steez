CREATE TABLE tb_test_accounts (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    account_no VARCHAR(64) UNIQUE NOT NULL COMMENT '账户号',
    property TINYINT DEFAULT 1 COMMENT '1-正常 0 未知 -1-不正常',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    INDEX idx_account_no (account_no),
    INDEX idx_property (property),
    INDEX idx_created_at (created_at),
    INDEX idx_updated_at (updated_at),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='账户表';