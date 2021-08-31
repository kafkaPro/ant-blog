package db

import (
	"context"
	"github.com/gogf/gf/database/gdb"

	"github.com/gogf/gf/os/gtime"

	"github.com/gogf/gf/frame/g"
)

func TableApis() (err error) {
	err = g.DB(Db).Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		_, e := tx.Ctx(ctx).Exec("DROP TABLE IF EXISTS `apis`;")
		if e != nil {
			return e
		}

		_, e = tx.Ctx(ctx).Exec("CREATE TABLE `apis` (  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',   `create_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',   `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',   `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',   `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT 'api路径',   `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT 'api中文描述',   `api_group` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT 'api组',   `method` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT 'POST' COMMENT '方法',   PRIMARY KEY (`id`) USING BTREE,   KEY `idx_apis_delete_at` (`delete_at`) USING BTREE ) ENGINE=InnoDB AUTO_INCREMENT=148 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;")
		if e != nil {
			return e
		}
		return nil
	})
	//_, err = g.DB(Db).Exec("DROP TABLE IF EXISTS `apis`; CREATE TABLE `apis` (  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',   `create_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',   `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',   `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',   `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT 'api路径',   `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT 'api中文描述',   `api_group` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT 'api组',   `method` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT 'POST' COMMENT '方法',   PRIMARY KEY (`id`) USING BTREE,   KEY `idx_apis_delete_at` (`delete_at`) USING BTREE ) ENGINE=InnoDB AUTO_INCREMENT=148 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;")
	return
}

func TableJwts() (err error) {
	err = g.DB(Db).Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		_, e := tx.Ctx(ctx).Exec("DROP TABLE IF EXISTS `jwts`; ")
		if e != nil {
			return e
		}

		_, e = tx.Ctx(ctx).Exec("CREATE TABLE `jwts` (   `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',   `create_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',   `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',   `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',   `jwt` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT 'jwt',   PRIMARY KEY (`id`) USING BTREE,   KEY `idx_jwts_deleted_at` (`delete_at`) USING BTREE ) ENGINE=InnoDB AUTO_INCREMENT=145 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;")
		if e != nil {
			return e
		}
		return nil
	})
	return
}

func TableMenus() (err error) {
	err = g.DB(Db).Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		_, e := tx.Ctx(ctx).Exec("DROP TABLE IF EXISTS `menus`; ")
		if e != nil {
			return e
		}

		_, e = tx.Ctx(ctx).Exec("CREATE TABLE `menus` (   `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',   `create_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',   `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',   `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',   `menu_level` int unsigned DEFAULT NULL COMMENT '菜单等级(预留字段)',   `parent_id` varchar(255) DEFAULT NULL COMMENT '父菜单ID',   `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '路由path',   `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '路由name',   `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',   `component` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '前端文件路径',   `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '菜单名',   `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '菜单图标',   `sort` int DEFAULT NULL COMMENT '排序标记',   `keep_alive` tinyint(1) DEFAULT NULL COMMENT '是否缓存',   `default_menu` tinyint(1) DEFAULT NULL COMMENT '是否是基础路由(开发中)',   PRIMARY KEY (`id`) USING BTREE,   KEY `idx_menus_deleted_at` (`delete_at`) USING BTREE ) ENGINE=InnoDB AUTO_INCREMENT=513 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;")
		if e != nil {
			return e
		}
		return nil
	})
	return
}

func TableFiles() (err error) {
	err = g.DB(Db).Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		_, e := tx.Ctx(ctx).Exec("DROP TABLE IF EXISTS `files`; ")
		if e != nil {
			return e
		}

		_, e = tx.Ctx(ctx).Exec("CREATE TABLE `files` (   `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',   `create_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',   `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',   `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',   `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '文件名',   `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '文件地址',   `tag` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '文件标签',   `key` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '编号',   PRIMARY KEY (`id`) USING BTREE,   KEY `idx_files_deleted_at` (`delete_at`) USING BTREE ) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;")
		if e != nil {
			return e
		}
		return nil
	})
	return
}

func TableAdmins() (err error) {
	err = g.DB(Db).Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		_, e := tx.Ctx(ctx).Exec("DROP TABLE IF EXISTS `admins`; ")
		if e != nil {
			return e
		}

		_, e = tx.Ctx(ctx).Exec("CREATE TABLE `admins` (   `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',   `create_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',   `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',   `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',   `uuid` varchar(255) DEFAULT NULL COMMENT '用户唯一标识UUID',   `nickname` varchar(255) DEFAULT 'QMPlusUser' COMMENT '用户昵称',   `header_img` varchar(255) DEFAULT 'http://www.henrongyi.top/avatar/lufu.jpg' COMMENT '用户头像',   `authority_id` varchar(255) DEFAULT '888' COMMENT '用户角色ID',   `username` varchar(255) DEFAULT NULL COMMENT '用户名',   `password` varchar(255) DEFAULT NULL COMMENT '用户登录密码',   PRIMARY KEY (`id`),   KEY `idx_admins_deleted_at` (`delete_at`) ) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;")
		if e != nil {
			return e
		}
		return nil
	})
	return
}

func TableCustomers() (err error) {
	err = g.DB(Db).Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		_, e := tx.Ctx(ctx).Exec("DROP TABLE IF EXISTS `customers`; ")
		if e != nil {
			return e
		}

		_, e = tx.Ctx(ctx).Exec("CREATE TABLE `customers` (   `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',   `create_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',   `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',   `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',   `customer_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '客户名',   `customer_phone_data` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '客户电话',   `sys_user_id` int unsigned DEFAULT NULL COMMENT '负责员工id',   `sys_user_authority_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '负责员工角色',   PRIMARY KEY (`id`) USING BTREE,   KEY `idx_exa_customers_deleted_at` (`delete_at`) USING BTREE ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=COMPACT;")
		if e != nil {
			return e
		}
		return nil
	})
	return
}

func TableCasbinRule() (err error) {
	err = g.DB(Db).Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		_, e := tx.Ctx(ctx).Exec("DROP TABLE IF EXISTS `casbin_rule`; ")
		if e != nil {
			return e
		}

		_, e = tx.Ctx(ctx).Exec("CREATE TABLE `casbin_rule` (   `ptype` varchar(10) DEFAULT NULL,   `v0` varchar(256) DEFAULT NULL,   `v1` varchar(256) DEFAULT NULL,   `v2` varchar(256) DEFAULT NULL,   `v3` varchar(256) DEFAULT NULL,   `v4` varchar(256) DEFAULT NULL,   `v5` varchar(256) DEFAULT NULL ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;")
		if e != nil {
			return e
		}
		return nil
	})
	return
}

func TableOperations() (err error) {
	err = g.DB(Db).Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		_, e := tx.Ctx(ctx).Exec("DROP TABLE IF EXISTS `operations`;")
		if e != nil {
			return e
		}

		_, e = tx.Ctx(ctx).Exec("CREATE TABLE `operations` (   `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',   `create_at` datetime DEFAULT NULL COMMENT '创建时间',   `update_at` datetime DEFAULT NULL COMMENT '更新时间',   `delete_at` datetime DEFAULT NULL COMMENT '删除时间',   `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '请求ip',   `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '请求方法',   `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '请求路由',   `status` int DEFAULT NULL COMMENT '状态',   `latency` bigint DEFAULT NULL COMMENT '延迟',   `agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '代理',   `error_message` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL COMMENT '报错信息',   `request` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '请求Body',   `user_id` int DEFAULT NULL COMMENT '用户id',   `response` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin COMMENT '响应Body',   PRIMARY KEY (`id`) USING BTREE,   KEY `idx_operations_deleted_at` (`delete_at`) USING BTREE ) ENGINE=InnoDB AUTO_INCREMENT=392 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin ROW_FORMAT=COMPACT;")
		if e != nil {
			return e
		}
		return nil
	})
	return
}

func TableParameters() (err error) {
	err = g.DB(Db).Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		_, e := tx.Ctx(ctx).Exec("DROP TABLE IF EXISTS `parameters`;")
		if e != nil {
			return e
		}

		_, e = tx.Ctx(ctx).Exec("CREATE TABLE `parameters`  (   `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT comment '自增ID',   `create_at` datetime(0) NULL DEFAULT NULL comment '创建时间',   `update_at` datetime(0) NULL DEFAULT NULL comment '更新时间',   `delete_at` datetime(0) NULL DEFAULT NULL comment '删除时间',   `base_menu_id` int(10) UNSIGNED NULL DEFAULT NULL comment 'BaseMenu的ID',   `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL comment '地址栏携带参数为params还是query',   `key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL comment '地址栏携带参数的key',   `value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL comment '地址栏携带参数的值',   PRIMARY KEY (`id`) USING BTREE,   INDEX `idx_sys_base_menu_parameters_deleted_at`(`delete_at`) USING BTREE ) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin  ROW_FORMAT = Compact;")
		if e != nil {
			return e
		}
		return nil
	})
	return
}

func TableAuthorities() (err error) {
	err = g.DB(Db).Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		_, e := tx.Ctx(ctx).Exec("DROP TABLE IF EXISTS `authorities`;")
		if e != nil {
			return e
		}

		_, e = tx.Ctx(ctx).Exec("CREATE TABLE `authorities` (   `authority_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色ID',   `authority_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '角色名',   `parent_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '父角色ID',   `create_at` datetime DEFAULT NULL COMMENT '创建时间',   `update_at` datetime DEFAULT NULL COMMENT '更新时间',   `delete_at` datetime DEFAULT NULL COMMENT '删除时间',   PRIMARY KEY (`authority_id`) USING BTREE,   UNIQUE KEY `authority_id` (`authority_id`) USING BTREE,   KEY `idx_sys_authorities_deleted_at` (`delete_at`) USING BTREE ) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;")
		if e != nil {
			return e
		}
		return nil
	})
	return
}

//func TableSimpleUpload() (err error) {
//	_, err = g.DB(Db).Exec("DROP TABLE IF EXISTS `authorities`; CREATE TABLE `authorities` (   `authority_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色ID',   `authority_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '角色名',   `parent_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '父角色ID',   `create_at` datetime DEFAULT NULL COMMENT '创建时间',   `update_at` datetime DEFAULT NULL COMMENT '更新时间',   `delete_at` datetime DEFAULT NULL COMMENT '删除时间',   PRIMARY KEY (`authority_id`) USING BTREE,   UNIQUE KEY `authority_id` (`authority_id`) USING BTREE,   KEY `idx_sys_authorities_deleted_at` (`delete_at`) USING BTREE ) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;")
//	return
//}

func TableDictionaries() (err error) {
	err = g.DB(Db).Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		_, e := tx.Ctx(ctx).Exec("DROP TABLE IF EXISTS `dictionaries`;")
		if e != nil {
			return e
		}

		_, e = tx.Ctx(ctx).Exec("CREATE TABLE `dictionaries` (   `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',   `create_at` datetime DEFAULT NULL COMMENT '创建时间',   `update_at` datetime DEFAULT NULL COMMENT '更新时间',   `delete_at` datetime DEFAULT NULL COMMENT '删除时间',   `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '字典名（中）',   `type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '字典名（英）',   `status` tinyint(1) DEFAULT NULL COMMENT '状态',   `desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '描述',   PRIMARY KEY (`id`) USING BTREE,   KEY `idx_sys_dictionaries_deleted_at` (`delete_at`) USING BTREE ) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;")
		if e != nil {
			return e
		}
		return nil
	})
	return
}

func TableAuthorityMenu() (err error) {
	err = g.DB(Db).Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		_, e := tx.Ctx(ctx).Exec("DROP TABLE IF EXISTS `authority_menu`;")
		if e != nil {
			return e
		}

		_, e = tx.Ctx(ctx).Exec("CREATE TABLE `authority_menu` (   `authority_id` varchar(255) NOT NULL COMMENT '权限id',   `menu_id` varchar(255) NOT NULL COMMENT '菜单id',   KEY `menu_id` (`menu_id`),   KEY `authority_id` (`authority_id`) ) ENGINE=InnoDB DEFAULT CHARSET=latin1;")
		if e != nil {
			return e
		}
		return nil
	})
	return
}

func TableBreakpointFiles() (err error) {
	err = g.DB(Db).Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		_, e := tx.Ctx(ctx).Exec("DROP TABLE IF EXISTS `breakpoint_files`;")
		if e != nil {
			return e
		}

		_, e = tx.Ctx(ctx).Exec("CREATE TABLE `breakpoint_files` (   `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',   `create_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',   `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',   `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',   `file_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '文件名',   `file_md5` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '文件md5',   `file_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '文件路径',   `chunk_id` int DEFAULT NULL COMMENT '关联id',   `chunk_total` int DEFAULT NULL COMMENT '切片总数',   `is_finish` tinyint(1) DEFAULT NULL COMMENT '是否完整',   PRIMARY KEY (`id`) USING BTREE,   KEY `idx_breakpoint_files_deleted_at` (`delete_at`) USING BTREE ) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;")
		if e != nil {
			return e
		}
		return nil
	})
	return
}

func TableBreakpointChucks() (err error) {
	err = g.DB(Db).Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		_, e := tx.Ctx(ctx).Exec("DROP TABLE IF EXISTS `breakpoint_chucks`;")
		if e != nil {
			return e
		}

		_, e = tx.Ctx(ctx).Exec("CREATE TABLE `breakpoint_chucks` (   `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',   `create_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',   `update_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',   `delete_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',   `exa_file_id` int unsigned DEFAULT NULL COMMENT '文件id',   `file_chunk_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '切片路径',   `file_chunk_number` int DEFAULT NULL COMMENT '切片标号',   PRIMARY KEY (`id`) USING BTREE,   KEY `idx_exa_file_chunks_deleted_at` (`delete_at`) USING BTREE ) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;")
		if e != nil {
			return e
		}
		return nil
	})
	return
}

func TableDictionaryDetails() (err error) {
	err = g.DB(Db).Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		_, e := tx.Ctx(ctx).Exec("DROP TABLE IF EXISTS `dictionary_details`;")
		if e != nil {
			return e
		}

		_, e = tx.Ctx(ctx).Exec("CREATE TABLE `dictionary_details` (   `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',   `create_at` datetime DEFAULT NULL COMMENT '创建时间',   `update_at` datetime DEFAULT NULL COMMENT '更新时间',   `delete_at` datetime DEFAULT NULL COMMENT '删除时间',   `label` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '展示值',   `value` int DEFAULT NULL COMMENT '字典值',   `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',   `sort` int DEFAULT NULL COMMENT '排序标记',   `dictionary_id` int DEFAULT NULL COMMENT '关联标记',   PRIMARY KEY (`id`) USING BTREE,   KEY `idx_sys_dictionary_details_deleted_at` (`delete_at`) USING BTREE ) ENGINE=InnoDB AUTO_INCREMENT=66 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;")
		if e != nil {
			return e
		}
		return nil
	})
	return
}

func TableAuthorityResources() (err error) {
	err = g.DB(Db).Transaction(context.TODO(), func(ctx context.Context, tx *gdb.TX) error {
		_, e := tx.Ctx(ctx).Exec("DROP TABLE IF EXISTS `authority_resources`;")
		if e != nil {
			return e
		}

		_, e = tx.Ctx(ctx).Exec("CREATE TABLE `authority_resources` (   `authority_id` varchar(255) NOT NULL COMMENT '权限id',   `resources_id` varchar(255) DEFAULT NULL COMMENT '资源id',   KEY `resources_id` (`resources_id`),   KEY `authority_id` (`authority_id`) ) ENGINE=InnoDB DEFAULT CHARSET=latin1;")
		if e != nil {
			return e
		}
		return nil
	})
	return
}

func DataApis() (err error) {
	var tx *gdb.TX
	if tx, err = g.DB(Db).Begin(); err != nil {
		panic(err)
	}
	db := g.DB(Db).Model("apis").Safe()
	_, err = db.Data(g.List{
		{"id": 1, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/base/login", "description": "用户登录", "api_group": "base", "method": "POST"},
		{"id": 2, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/base/register", "description": "用户注册", "api_group": "base", "method": "POST"},
		{"id": 3, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/api/createApi", "description": "创建api", "api_group": "api", "method": "POST"},
		{"id": 4, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/api/getApiList", "description": "获取api列表", "api_group": "api", "method": "POST"},
		{"id": 5, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/api/getApiById", "description": "获取api详细信息", "api_group": "api", "method": "POST"},
		{"id": 6, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/api/deleteApi", "description": "删除Api", "api_group": "api", "method": "POST"},
		{"id": 7, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/api/updateApi", "description": "更新Api", "api_group": "api", "method": "POST"},
		{"id": 8, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/api/getAllApis", "description": "获取所有api", "api_group": "api", "method": "POST"},
		{"id": 9, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/authority/createAuthority", "description": "创建角色", "api_group": "authority", "method": "POST"},
		{"id": 10, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/authority/deleteAuthority", "description": "删除角色", "api_group": "authority", "method": "POST"},
		{"id": 11, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/authority/getAuthorityList", "description": "获取角色列表", "api_group": "authority", "method": "POST"},
		{"id": 12, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/menu/getMenu", "description": "获取菜单树", "api_group": "menu", "method": "POST"},
		{"id": 13, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/menu/getMenuList", "description": "分页获取基础menu列表", "api_group": "menu", "method": "POST"},
		{"id": 14, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/menu/addBaseMenu", "description": "新增菜单", "api_group": "menu", "method": "POST"},
		{"id": 15, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/menu/getBaseMenuTree", "description": "获取用户动态路由", "api_group": "menu", "method": "POST"},
		{"id": 16, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/menu/addMenuAuthority", "description": "增加menu和角色关联关系", "api_group": "menu", "method": "POST"},
		{"id": 17, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/menu/getMenuAuthority", "description": "获取指定角色menu", "api_group": "menu", "method": "POST"},
		{"id": 18, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/menu/deleteBaseMenu", "description": "删除菜单", "api_group": "menu", "method": "POST"},
		{"id": 19, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/menu/updateBaseMenu", "description": "更新菜单", "api_group": "menu", "method": "POST"},
		{"id": 20, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/menu/getBaseMenuById", "description": "根据id获取菜单", "api_group": "menu", "method": "POST"},
		{"id": 21, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/user/changePassword", "description": "修改密码", "api_group": "user", "method": "POST"},
		{"id": 22, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/user/uploadHeaderImg", "description": "上传头像", "api_group": "user", "method": "POST"},
		{"id": 23, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/user/getInfoList", "description": "分页获取用户列表", "api_group": "user", "method": "POST"},
		{"id": 24, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/user/getUserList", "description": "获取用户列表", "api_group": "user", "method": "POST"},
		{"id": 25, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/user/setUserAuthority", "description": "修改用户角色", "api_group": "user", "method": "POST"},
		{"id": 26, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/fileUploadAndDownload/upload", "description": "文件上传示例", "api_group": "fileUploadAndDownload", "method": "POST"},
		{"id": 27, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/fileUploadAndDownload/getFileList", "description": "获取上传文件列表", "api_group": "fileUploadAndDownload", "method": "POST"},
		{"id": 28, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/casbin/updateCasbin", "description": "更改角色api权限", "api_group": "casbin", "method": "POST"},
		{"id": 29, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/casbin/getPolicyPathByAuthorityId", "description": "获取权限列表", "api_group": "casbin", "method": "POST"},
		{"id": 30, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/fileUploadAndDownload/deleteFile", "description": "删除文件", "api_group": "fileUploadAndDownload", "method": "POST"},
		{"id": 31, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/jwt/jsonInBlacklist", "description": "jwt加入黑名单", "api_group": "jwt", "method": "POST"},
		{"id": 32, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/authority/setDataAuthority", "description": "设置角色资源权限", "api_group": "authority", "method": "POST"},
		{"id": 33, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/system/getSystemConfig", "description": "获取配置文件内容", "api_group": "system", "method": "POST"},
		{"id": 34, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/system/setSystemConfig", "description": "设置配置文件内容", "api_group": "system", "method": "POST"},
		{"id": 35, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/customer/customer", "description": "创建客户", "api_group": "customer", "method": "POST"},
		{"id": 36, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/customer/customer", "description": "更新客户", "api_group": "customer", "method": "PUT"},
		{"id": 37, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/customer/customer", "description": "删除客户", "api_group": "customer", "method": "DELETE"},
		{"id": 38, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/customer/customer", "description": "获取单一客户", "api_group": "customer", "method": "GET"},
		{"id": 39, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/customer/customerList", "description": "获取客户列表", "api_group": "customer", "method": "GET"},
		{"id": 40, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/casbin/casbinTest/:pathParam", "description": "RESTFUL模式测试", "api_group": "casbin", "method": "GET"},
		{"id": 41, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/autoCode/createTemp", "description": "自动化代码", "api_group": "autoCode", "method": "POST"},
		{"id": 42, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/authority/updateAuthority", "description": "更新角色信息", "api_group": "authority", "method": "PUT"},
		{"id": 43, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/authority/copyAuthority", "description": "拷贝角色", "api_group": "authority", "method": "POST"},
		{"id": 44, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/user/deleteUser", "description": "删除用户", "api_group": "user", "method": "DELETE"},
		{"id": 45, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/sysDictionaryDetail/createSysDictionaryDetail", "description": "新增字典内容", "api_group": "sysDictionaryDetail", "method": "POST"},
		{"id": 46, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/sysDictionaryDetail/deleteSysDictionaryDetail", "description": "删除字典内容", "api_group": "sysDictionaryDetail", "method": "DELETE"},
		{"id": 47, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/sysDictionaryDetail/updateSysDictionaryDetail", "description": "更新字典内容", "api_group": "sysDictionaryDetail", "method": "PUT"},
		{"id": 48, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/sysDictionaryDetail/findSysDictionaryDetail", "description": "根据ID获取字典内容", "api_group": "sysDictionaryDetail", "method": "GET"},
		{"id": 49, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/sysDictionaryDetail/getSysDictionaryDetailList", "description": "获取字典内容列表", "api_group": "sysDictionaryDetail", "method": "GET"},
		{"id": 50, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/sysDictionary/createSysDictionary", "description": "新增字典", "api_group": "sysDictionary", "method": "POST"},
		{"id": 51, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/sysDictionary/deleteSysDictionary", "description": "删除字典", "api_group": "sysDictionary", "method": "DELETE"},
		{"id": 52, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/sysDictionary/updateSysDictionary", "description": "更新字典", "api_group": "sysDictionary", "method": "PUT"},
		{"id": 53, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/sysDictionary/findSysDictionary", "description": "根据ID获取字典", "api_group": "sysDictionary", "method": "GET"},
		{"id": 54, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/sysDictionary/getSysDictionaryList", "description": "获取字典列表", "api_group": "sysDictionary", "method": "GET"},
		{"id": 55, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/autoCode/getTables", "description": "获取数据库表", "api_group": "autoCode", "method": "GET"},
		{"id": 56, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/autoCode/getDB", "description": "获取所有数据库", "api_group": "autoCode", "method": "GET"},
		{"id": 57, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/autoCode/getColume", "description": "获取所选table的所有字段", "api_group": "autoCode", "method": "GET"},
		{"id": 58, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/user/setUserInfo", "description": "设置用户信息", "api_group": "user", "method": "PUT"},
		{"id": 59, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/simpleUploader/upload", "description": "插件版分片上传", "api_group": "simpleUploader", "method": "POST"},
		{"id": 60, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/simpleUploader/checkFileMd5", "description": "文件完整度验证", "api_group": "simpleUploader", "method": "GET"},
		{"id": 61, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "path": "/simpleUploader/mergeFileMd5", "description": "上传完成合并文件", "api_group": "simpleUploader", "method": "GET"},
	}).Batch(10).Insert()
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}

func DataFiles() (err error) {
	var tx *gdb.TX
	if tx, err = g.DB(Db).Begin(); err != nil {
		panic(err)
	}
	db := g.DB(Db).Model("files").Safe()
	_, err = db.Data(g.List{
		{"id": 1, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "name": "10.png", "url": "http://qmplusimg.henrongyi.top/gvalogo.png", "tag": "png", "key": "158787308910.png"},
		{"id": 2, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "name": "logo.png", "url": "http://qmplusimg.henrongyi.top/1576554439myAvatar.png", "tag": "png", "key": "1587973709logo.png"},
	}).Batch(10).Insert()
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}

func DataMenus() (err error) {
	var tx *gdb.TX
	if tx, err = g.DB(Db).Begin(); err != nil {
		panic(err)
	}
	db := g.DB(Db).Model("menus").Safe()
	_, err = db.Data(g.List{
		{"id": 1, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 0, "name": "dashboard", "path": "dashboard", "hidden": 0, "component": "view/dashboard/index.vue", "title": "仪表盘", "icon": "setting", "sort": 1, "keep_alive": 0, "default_menu": 0},
		{"id": 2, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 0, "name": "about", "path": "about", "hidden": 0, "component": "view/about/index.vue", "title": "关于我们", "icon": "info", "sort": 7, "keep_alive": 0, "default_menu": 0},
		{"id": 3, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 0, "name": "admin", "path": "superAdmin", "hidden": 0, "component": "view/superAdmin/index.vue", "title": "超级管理员", "icon": "user-solid", "sort": 3, "keep_alive": 0, "default_menu": 0},
		{"id": 4, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 3, "name": "authority", "path": "authority", "hidden": 0, "component": "view/superAdmin/authority/authority.vue", "title": "角色管理", "icon": "s-custom", "sort": 1, "keep_alive": 0, "default_menu": 0},
		{"id": 5, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 3, "name": "menu", "path": "menu", "hidden": 0, "component": "view/superAdmin/menu/menu.vue", "title": "菜单管理", "icon": "s-order", "sort": 2, "keep_alive": 1, "default_menu": 0},
		{"id": 6, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 3, "name": "api", "path": "api", "hidden": 0, "component": "view/superAdmin/api/api.vue", "title": "api管理", "icon": "s-platform", "sort": 3, "keep_alive": 1, "default_menu": 0},
		{"id": 7, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 3, "name": "user", "path": "user", "hidden": 0, "component": "view/superAdmin/user/user.vue", "title": "用户管理", "icon": "coordinate", "sort": 4, "keep_alive": 0, "default_menu": 0},
		{"id": 8, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 0, "name": "person", "path": "person", "hidden": 1, "component": "view/person/person.vue", "title": "个人信息", "icon": "message-solid", "sort": 4, "keep_alive": 0, "default_menu": 0},
		{"id": 9, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 0, "name": "example", "path": "example", "hidden": 0, "component": "view/example/index.vue", "title": "示例文件", "icon": "s-management", "sort": 6, "keep_alive": 0, "default_menu": 0},
		{"id": 10, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 9, "name": "table", "path": "table", "hidden": 0, "component": "view/example/table/table.vue", "title": "表格示例", "icon": "s-order", "sort": 1, "keep_alive": 0, "default_menu": 0},
		{"id": 11, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 9, "name": "form", "path": "form", "hidden": 0, "component": "view/example/form/form.vue", "title": "表单示例", "icon": "document", "sort": 2, "keep_alive": 0, "default_menu": 0},
		{"id": 12, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 9, "name": "rte", "path": "rte", "hidden": 0, "component": "view/example/rte/rte.vue", "title": "富文本编辑器", "icon": "reading", "sort": 3, "keep_alive": 0, "default_menu": 0},
		{"id": 13, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 9, "name": "excel", "path": "excel", "hidden": 0, "component": "view/example/excel/excel.vue", "title": "excel导入导出", "icon": "s-marketing", "sort": 4, "keep_alive": 0, "default_menu": 0},
		{"id": 14, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 9, "name": "upload", "path": "upload", "hidden": 0, "component": "view/example/upload/upload.vue", "title": "上传下载", "icon": "upload", "sort": 5, "keep_alive": 0, "default_menu": 0},
		{"id": 15, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 9, "name": "breakpoint", "path": "breakpoint", "hidden": 0, "component": "view/example/breakpoint/breakpoint.vue", "title": "断点续传", "icon": "upload", "sort": 6, "keep_alive": 0, "default_menu": 0},
		{"id": 16, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 9, "name": "customer", "path": "customer", "hidden": 0, "component": "view/example/customer/customer.vue", "title": "客户列表（资源示例）", "icon": "s-custom", "sort": 7, "keep_alive": 0, "default_menu": 0},
		{"id": 17, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 0, "name": "systemTools", "path": "systemTools", "hidden": 0, "component": "view/systemTools/index.vue", "title": "系统工具", "icon": "s-cooperation", "sort": 5, "keep_alive": 0, "default_menu": 0},
		{"id": 18, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 17, "name": "autoCode", "path": "autoCode", "hidden": 0, "component": "view/systemTools/autoCode/index.vue", "title": "代码生成器", "icon": "cpu", "sort": 1, "keep_alive": 1, "default_menu": 0},
		{"id": 19, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 17, "name": "formCreate", "path": "formCreate", "hidden": 0, "component": "view/systemTools/formCreate/index.vue", "title": "表单生成器", "icon": "magic-stick", "sort": 2, "keep_alive": 1, "default_menu": 0},
		{"id": 20, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 17, "name": "system", "path": "system", "hidden": 0, "component": "view/systemTools/system/system.vue", "title": "系统配置", "icon": "s-operation", "sort": 3, "keep_alive": 0, "default_menu": 0},
		{"id": 21, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 0, "name": "iconList", "path": "iconList", "hidden": 0, "component": "view/iconList/index.vue", "title": "图标集合", "icon": "star-on", "sort": 2, "keep_alive": 0, "default_menu": 0},
		{"id": 22, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 3, "name": "dictionary", "path": "dictionary", "hidden": 0, "component": "view/superAdmin/dictionary/sysDictionary.vue", "title": "字典管理", "icon": "notebook-2", "sort": 5, "keep_alive": 0, "default_menu": 0},
		{"id": 23, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 3, "name": "dictionaryDetail", "path": "dictionaryDetail/:id", "hidden": 1, "component": "view/superAdmin/dictionary/sysDictionaryDetail.vue", "title": "字典详情", "icon": "s-order", "sort": 1, "keep_alive": 0, "default_menu": 0},
		{"id": 24, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 3, "name": "operation", "path": "operation", "hidden": 0, "component": "view/superAdmin/operation/sysOperationRecord.vue", "title": "操作历史", "icon": "time", "sort": 6, "keep_alive": 0, "default_menu": 0},
		{"id": 25, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 0, "name": "https://www.gin-vue-admin.com", "path": "https://www.gin-vue-admin.com", "hidden": 0, "component": "/", "title": "官方网站", "icon": "s-home", "sort": 0, "keep_alive": 0, "default_menu": 0},
		{"id": 26, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "menu_level": 0, "parent_id": 9, "name": "simpleUploader", "path": "simpleUploader", "hidden": 0, "component": "view/example/simpleUploader/simpleUploader", "title": "断点续传（插件版）", "icon": "upload", "sort": 6, "keep_alive": 0, "default_menu": 0},
	}).Batch(10).Insert()
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}

func DataAdmins() (err error) {
	var tx *gdb.TX
	if tx, err = g.DB(Db).Begin(); err != nil {
		panic(err)
	}
	db := g.DB(Db).Model("admins").Safe()
	_, err = db.Data(g.List{
		{"id": 1, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "uuid": "b4c54e5a-d015-4f8c-9f01-624c527a63ae", "nickname": "超级管理员", "header_img": "http://qmplusimg.henrongyi.top/1571627762timg.jpg", "authority_id": "888", "username": "admin", "password": "$2a$10$zF5PNCWobve/0RBk.3k03eAGwyvDevFBFd3AZUwETjMhYpZwNooba"},
		{"id": 2, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "uuid": "fd6ef79b-944c-4888-8377-abe2d2608858", "nickname": "QMPlusUser", "header_img": "http://qmplusimg.henrongyi.top/1572075907logo.png", "authority_id": "9528", "username": "a303176530", "password": "$2a$10$zF5PNCWobve/0RBk.3k03eAGwyvDevFBFd3AZUwETjMhYpZwNooba"},
	}).Batch(10).Insert()
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}

func DataCasbinRule() (err error) {
	var tx *gdb.TX
	if tx, err = g.DB(Db).Begin(); err != nil {
		panic(err)
	}
	db := g.DB(Db).Model("casbin_rule").Safe()
	_, err = db.Data(g.List{
		{"ptype": "p", "v0": "888", "v1": "/base/login", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/base/register", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/api/createApi", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/api/getApiList", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/api/getApiById", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/api/deleteApi", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/api/updateApi", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/api/getAllApis", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/authority/createAuthority", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/authority/deleteAuthority", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/authority/getAuthorityList", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/authority/setDataAuthority", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/authority/updateAuthority", "v2": "PUT"},
		{"ptype": "p", "v0": "888", "v1": "/authority/copyAuthority", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/menu/getMenu", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/menu/getMenuList", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/menu/addBaseMenu", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/menu/getBaseMenuTree", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/menu/addMenuAuthority", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/menu/getMenuAuthority", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/menu/deleteBaseMenu", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/menu/updateBaseMenu", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/menu/getBaseMenuById", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/user/changePassword", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/user/uploadHeaderImg", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/user/getInfoList", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/user/getUserList", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/user/setUserAuthority", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/user/deleteUser", "v2": "DELETE"},
		{"ptype": "p", "v0": "888", "v1": "/user/setUserInfo", "v2": "PUT"},
		{"ptype": "p", "v0": "888", "v1": "/fileUploadAndDownload/upload", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/fileUploadAndDownload/getFileList", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/fileUploadAndDownload/deleteFile", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/casbin/updateCasbin", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/casbin/getPolicyPathByAuthorityId", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/casbin/casbinTest/:pathParam", "v2": "GET"},
		{"ptype": "p", "v0": "888", "v1": "/jwt/jsonInBlacklist", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/system/getSystemConfig", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/system/setSystemConfig", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/customer/customer", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/customer/customer", "v2": "PUT"},
		{"ptype": "p", "v0": "888", "v1": "/customer/customer", "v2": "DELETE"},
		{"ptype": "p", "v0": "888", "v1": "/customer/customer", "v2": "GET"},
		{"ptype": "p", "v0": "888", "v1": "/customer/customerList", "v2": "GET"},
		{"ptype": "p", "v0": "888", "v1": "/autoCode/createTemp", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/autoCode/getTables", "v2": "GET"},
		{"ptype": "p", "v0": "888", "v1": "/autoCode/getDB", "v2": "GET"},
		{"ptype": "p", "v0": "888", "v1": "/autoCode/getColume", "v2": "GET"},
		{"ptype": "p", "v0": "888", "v1": "/sysDictionaryDetail/createSysDictionaryDetail", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/sysDictionaryDetail/deleteSysDictionaryDetail", "v2": "DELETE"},
		{"ptype": "p", "v0": "888", "v1": "/sysDictionaryDetail/updateSysDictionaryDetail", "v2": "PUT"},
		{"ptype": "p", "v0": "888", "v1": "/sysDictionaryDetail/findSysDictionaryDetail", "v2": "GET"},
		{"ptype": "p", "v0": "888", "v1": "/sysDictionaryDetail/getSysDictionaryDetailList", "v2": "GET"},
		{"ptype": "p", "v0": "888", "v1": "/sysDictionary/createSysDictionary", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/sysDictionary/deleteSysDictionary", "v2": "DELETE"},
		{"ptype": "p", "v0": "888", "v1": "/sysDictionary/updateSysDictionary", "v2": "PUT"},
		{"ptype": "p", "v0": "888", "v1": "/sysDictionary/findSysDictionary", "v2": "GET"},
		{"ptype": "p", "v0": "888", "v1": "/sysDictionary/getSysDictionaryList", "v2": "GET"},
		{"ptype": "p", "v0": "888", "v1": "/sysOperationRecord/createSysOperationRecord", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/sysOperationRecord/deleteSysOperationRecord", "v2": "DELETE"},
		{"ptype": "p", "v0": "888", "v1": "/sysOperationRecord/updateSysOperationRecord", "v2": "PUT"},
		{"ptype": "p", "v0": "888", "v1": "/sysOperationRecord/findSysOperationRecord", "v2": "GET"},
		{"ptype": "p", "v0": "888", "v1": "/sysOperationRecord/getSysOperationRecordList", "v2": "GET"},
		{"ptype": "p", "v0": "888", "v1": "/sysOperationRecord/deleteSysOperationRecordByIds", "v2": "DELETE"},
		{"ptype": "p", "v0": "888", "v1": "/simpleUploader/upload", "v2": "POST"},
		{"ptype": "p", "v0": "888", "v1": "/simpleUploader/checkFileMd5", "v2": "GET"},
		{"ptype": "p", "v0": "888", "v1": "/simpleUploader/mergeFileMd5", "v2": "GET"},
	}).Batch(100).Insert()
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}

func DataAuthorities() (err error) {
	var tx *gdb.TX
	if tx, err = g.DB(Db).Begin(); err != nil {
		panic(err)
	}
	db := g.DB(Db).Model("authorities").Safe()
	_, err = db.Data(g.List{ // 插入数据
		{"authority_id": "888", "authority_name": "超级管理员", "parent_id": "0", "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil},
		{"authority_id": "8881", "authority_name": "超级管理员子账号", "parent_id": "888", "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil},
		{"authority_id": "9528", "authority_name": "测试角色", "parent_id": "0", "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil},
	}).Batch(10).Insert()
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}

func DataDictionaries() (err error) {
	var tx *gdb.TX
	if tx, err = g.DB(Db).Begin(); err != nil {
		panic(err)
	}
	db := g.DB(Db).Model("dictionaries").Safe()
	_, err = db.Data(g.List{
		{"id": 1, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "name": "性别", "type": "sex", "status": 1, "desc": "性别字典"},
		{"id": 2, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "name": "数据库int类型", "type": "int", "status": 1, "desc": "数据库int类型"},
		{"id": 3, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "name": "数据库时间日期类型", "type": "time.Time", "status": 1, "desc": "数据库时间日期类型"},
		{"id": 4, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "name": "数据库浮点型", "type": "float64", "status": 1, "desc": "数据库浮点型"},
		{"id": 5, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "name": "数据库字符串", "type": "string", "status": 1, "desc": "数据库字符串"},
		{"id": 6, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "name": "数据库bool类型", "type": "bool", "status": 1, "desc": "数据库bool类型"},
	}).Batch(10).Insert()
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}

func DataAuthorityMenus() (err error) {
	var tx *gdb.TX
	if tx, err = g.DB(Db).Begin(); err != nil {
		panic(err)
	}
	db := g.DB(Db).Model("authority_menu").Safe()
	_, err = db.Data(g.List{ // 插入数据
		{"authority_id": "888", "menu_id": "1"},
		{"authority_id": "888", "menu_id": "2"},
		{"authority_id": "888", "menu_id": "3"},
		{"authority_id": "888", "menu_id": "4"},
		{"authority_id": "888", "menu_id": "5"},
		{"authority_id": "888", "menu_id": "6"},
		{"authority_id": "888", "menu_id": "7"},
		{"authority_id": "888", "menu_id": "8"},
		{"authority_id": "888", "menu_id": "9"},
		{"authority_id": "888", "menu_id": "10"},
		{"authority_id": "888", "menu_id": "11"},
		{"authority_id": "888", "menu_id": "12"},
		{"authority_id": "888", "menu_id": "13"},
		{"authority_id": "888", "menu_id": "14"},
		{"authority_id": "888", "menu_id": "15"},
		{"authority_id": "888", "menu_id": "16"},
		{"authority_id": "888", "menu_id": "17"},
		{"authority_id": "888", "menu_id": "18"},
		{"authority_id": "888", "menu_id": "19"},
		{"authority_id": "888", "menu_id": "20"},
		{"authority_id": "888", "menu_id": "21"},
		{"authority_id": "888", "menu_id": "22"},
		{"authority_id": "888", "menu_id": "23"},
		{"authority_id": "888", "menu_id": "24"},
		{"authority_id": "888", "menu_id": "25"},
		{"authority_id": "888", "menu_id": "26"},
	}).Batch(10).Insert()
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}

func DataDictionaryDetails() (err error) {
	var tx *gdb.TX
	if tx, err = g.DB(Db).Begin(); err != nil {
		panic(err)
	}
	db := g.DB(Db).Model("dictionary_details").Safe()
	_, err = db.Data(g.List{
		{"id": 1, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "smallint", "value": 1, "status": 1, "sort": 1, "dictionary_id": 2},
		{"id": 2, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "mediumint", "value": 2, "status": 1, "sort": 2, "dictionary_id": 2},
		{"id": 3, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "int", "value": 3, "status": 1, "sort": 3, "dictionary_id": 2},
		{"id": 4, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "bigint", "value": 4, "status": 1, "sort": 4, "dictionary_id": 2},
		{"id": 5, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "data", "value": 0, "status": 1, "sort": 0, "dictionary_id": 3},
		{"id": 6, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "time", "value": 1, "status": 1, "sort": 1, "dictionary_id": 3},
		{"id": 7, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "year", "value": 2, "status": 1, "sort": 2, "dictionary_id": 3},
		{"id": 8, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "datetime", "value": 3, "status": 1, "sort": 3, "dictionary_id": 3},
		{"id": 9, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "timestamp", "value": 4, "status": 1, "sort": 4, "dictionary_id": 3},
		{"id": 10, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "float", "value": 0, "status": 1, "sort": 0, "dictionary_id": 4},
		{"id": 11, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "double", "value": 0, "status": 1, "sort": 0, "dictionary_id": 4},
		{"id": 12, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "decimal", "value": 0, "status": 1, "sort": 0, "dictionary_id": 4},
		{"id": 13, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "char", "value": 0, "status": 1, "sort": 0, "dictionary_id": 5},
		{"id": 14, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "varchar", "value": 0, "status": 1, "sort": 0, "dictionary_id": 5},
		{"id": 15, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "tinyblob", "value": 0, "status": 1, "sort": 0, "dictionary_id": 5},
		{"id": 16, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "tinytext", "value": 0, "status": 1, "sort": 0, "dictionary_id": 5},
		{"id": 17, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "text", "value": 0, "status": 1, "sort": 0, "dictionary_id": 5},
		{"id": 18, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "blob", "value": 0, "status": 1, "sort": 0, "dictionary_id": 5},
		{"id": 19, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "mediumblob", "value": 0, "status": 1, "sort": 0, "dictionary_id": 5},
		{"id": 20, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "mediumtext", "value": 0, "status": 1, "sort": 0, "dictionary_id": 5},
		{"id": 21, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "longblob", "value": 0, "status": 1, "sort": 0, "dictionary_id": 5},
		{"id": 22, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "longtext", "value": 0, "status": 1, "sort": 0, "dictionary_id": 5},
		{"id": 23, "create_at": gtime.Now(), "update_at": gtime.Now(), "delete_at": nil, "label": "tinyint", "value": 4, "status": 1, "sort": 4, "dictionary_id": 6},
	}).Batch(10).Insert()
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}
