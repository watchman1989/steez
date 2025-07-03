CREATE TABLE tb_test_transfer_records (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    transfer_no VARCHAR(64) UNIQUE NOT NULL COMMENT '转账流水号',
    from_account_id BIGINT NOT NULL COMMENT '转出账户ID',
    to_account_id BIGINT NOT NULL COMMENT '转入账户ID',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    INDEX idx_from_account (from_account_id),
    INDEX idx_to_account (to_account_id),
    INDEX idx_transfer_no (transfer_no),
    INDEX idx_created_at (created_at),
    INDEX idx_deleted_at (deleted_at),
    INDEX idx_updated_at (updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='转账记录表';