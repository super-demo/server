-- +goose Up
-- +goose StatementBegin
CREATE TABLE sites (
    site_id SERIAL PRIMARY KEY,
    site_type_id INT,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    short_description TEXT,
    url TEXT,
    image_url TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by INTEGER NOT NULL,
    deleted_at TIMESTAMP,

    FOREIGN KEY (site_type_id) REFERENCES site_types(site_type_id),
    FOREIGN KEY (created_by) REFERENCES users(user_id),
    FOREIGN KEY (updated_by) REFERENCES users(user_id)
);

-- Seed data
INSERT INTO sites (site_type_id, name, description, short_description, url, image_url, created_by, updated_by) 
VALUES 
    (2, 'Kasetsart University', 'Kasetsart University official site', 'Kasetsart University', 'https://www.ku.ac.th', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQh8BdSgouWizJ6JwSIEXHeFG8C5gAl4X1IuA&s', 1, 1),
    (1, 'Bangkhen Campus', 'Main campus of Kasetsart University', 'Bangkhen Campus', 'https://www.ku.ac.th/bangkhen', 'https://scontent.fbkk13-2.fna.fbcdn.net/v/t39.30808-6/305203334_479117977555430_4584347497113337500_n.jpg?_nc_cat=111&ccb=1-7&_nc_sid=6ee11a&_nc_ohc=gmarcBoVEuYQ7kNvgFDxrsR&_nc_oc=AdkUTUVIc5lbby2W60g7zJFuDGHAengHhEy_qXjBAESvE5UWJy4rwvVEQl7-ZEM6LHMZFKSqpJf-C8shVs_xthE4&_nc_zt=23&_nc_ht=scontent.fbkk13-2.fna&_nc_gid=K79NgbM5_iLxlnQcxHCbjA&oh=00_AYH-szSavNSajbEfVsd1I77jRvCkd-JSUoN8K4bvTkf5Mw&oe=67F01F65', 1, 1),
    (1, 'Kamphaeng Saen Campus', 'Agricultural campus of Kasetsart University', 'Kamphaeng Saen Campus', 'https://www.ku.ac.th/kps', 'https://scontent.fbkk13-1.fna.fbcdn.net/v/t1.6435-9/68765705_3026784780697566_2720886963409256448_n.jpg?_nc_cat=105&ccb=1-7&_nc_sid=cc71e4&_nc_ohc=DuFVka3dfFAQ7kNvgG2qJNR&_nc_oc=AdkVs-VvMZbWKamYJvBxqhZVQTZgTCaxoW5NYSrIyDBIoTbT2yjyDS4kzcPVYJusBZyV8rCuFvw4POa--3PBkM6l&_nc_zt=23&_nc_ht=scontent.fbkk13-1.fna&_nc_gid=tax5NMisKIWSun7nfJGedA&oh=00_AYHCsuXhA5YWWMwn8IvW5h3hGl1zCXoA-pFWHBsr8Pupuw&oe=6811C91F', 1, 1),
    (1, 'Sriracha Campus', 'Technology and business-focused campus of Kasetsart University', 'Sriracha Campus', 'https://www.ku.ac.th/sriracha', 'https://pbs.twimg.com/profile_images/1258304486780493827/aMSNrVWf_400x400.jpg', 1, 1),
    (1, 'Science', 'Science Faculty workspace', 'Science Faculty', 'https://science.ku.ac.th', 'https://scontent.fbkk12-1.fna.fbcdn.net/v/t39.30808-6/420854471_876307160957226_1354125451430269397_n.jpg?_nc_cat=101&ccb=1-7&_nc_sid=6ee11a&_nc_ohc=5uTHSkzYJ7YQ7kNvgG-OPkh&_nc_oc=Adm__6txQu5XivF0u8uzasRSyeLFHVjIzXYL_wLjs2GNsMzQi42sOwWcdHy7XHp5UpfDcBLe5n8GaJF3MBBvEGNQ&_nc_zt=23&_nc_ht=scontent.fbkk12-1.fna&_nc_gid=94jJ27OwIp99rWtPUzLfbA&oh=00_AYEvFq9GyJKT0surewZkV4PZr0hjrxWbWINC5qZWlBirPQ&oe=67EFF7BC', 1, 1),
    (1, 'Engineer', 'Engineering Faculty workspace', 'Engineering Faculty', 'https://engineering.ku.ac.th', 'https://scontent.fbkk12-2.fna.fbcdn.net/v/t39.30808-6/450612741_122163895064153111_4234078652444491256_n.jpg?_nc_cat=104&ccb=1-7&_nc_sid=6ee11a&_nc_ohc=APL8C2YDKVEQ7kNvgFVd9EY&_nc_oc=Admbko-lxNOxT9zs6SBzz5hLJaR-npyl_lBDwo_WZt5zeaznaEqSqBfTCPZR2eh7b83-Ou9v9FUEMJnaOCTTip2N&_nc_zt=23&_nc_ht=scontent.fbkk12-2.fna&_nc_gid=rtxDQ_QNbpcP7dUqQLD63g&oh=00_AYE2dcAGHf4HCthOxkQpYT9kVQGYQzv1juPvCFrqlNOtBg&oe=67F01589', 1, 1),
    (1, 'Agriculture', 'Agriculture Faculty workspace', 'Agriculture Faculty', 'https://agriculture.ku.ac.th', 'https://scontent.fbkk13-3.fna.fbcdn.net/v/t39.30808-6/310990465_180789204462063_2900171535943421022_n.jpg?_nc_cat=108&ccb=1-7&_nc_sid=6ee11a&_nc_ohc=iR_dmP3XraAQ7kNvgGXM9hu&_nc_oc=AdmmELqC-WMwygF6qyyjAZrajZJN1ylIlFTo1CHrMAaadIr9O5Govxx_7osgXYbgFpOQmLNob6zFx1r8_3AyctHf&_nc_zt=23&_nc_ht=scontent.fbkk13-3.fna&_nc_gid=SXXNoqLItnoQGvAgzKRZNg&oh=00_AYFP9KLDMNwJVZ2UfIy70cyZeUpT3lkhWRD_A3dToTrxPg&oe=67F02A28', 1, 1),
    (1, 'Economics', 'Economics Faculty workspace', 'Economics Faculty', 'https://economics.ku.ac.th', 'https://upload.wikimedia.org/wikipedia/commons/2/28/Logo_of_Kasetsart_Business_School%2C_Kasetsart_University.svg', 1, 1),
    (1, 'Fisheries', 'Fisheries Faculty workspace', 'Fisheries Faculty', 'https://fisheries.ku.ac.th', 'https://fish.ku.ac.th/sites/default/files/Management.jpg', 1, 1),
    (1, 'Forestry', 'Forestry Faculty workspace', 'Forestry Faculty', 'https://forestry.ku.ac.th', 'https://scontent.fbkk8-2.fna.fbcdn.net/v/t39.30808-6/282236236_5144941375584897_222690019403415550_n.jpg?_nc_cat=110&ccb=1-7&_nc_sid=6ee11a&_nc_ohc=MIzmIbLkcjkQ7kNvgFsyndY&_nc_oc=AdmRyFFAsy5LsftxtvbXm8KI_EKY-PejTFKcEWYbTq5F3supJ_YqrKDwuVYjtUIsVk3m1_KwjsXRK5XiUgHyHCD9&_nc_zt=23&_nc_ht=scontent.fbkk8-2.fna&_nc_gid=GkYlcCFp6mD3F9vYMB2e3A&oh=00_AYEqxn0orxtcW4NCscQ32O532nzLEQ588Sc3jeE-giP7HQ&oe=67F006AC', 1, 1),
    (1, 'Computer Science', 'Computer Science Department under Science Faculty', 'Computer Science', 'https://science.ku.ac.th/computer-science', 'https://scontent.fbkk13-3.fna.fbcdn.net/v/t39.30808-6/370139593_786202983515975_5138980421154670163_n.jpg?_nc_cat=108&ccb=1-7&_nc_sid=6ee11a&_nc_ohc=5HmKF9lAsA8Q7kNvgGNpjAh&_nc_oc=AdlblP4IILQyXgpCpPrpOE6IV5kW3yIpd4x7PQZzg5KzHnb9mX1si_KN1eUtUAuMiw-rPvLcyhDgBUwdVGQQ7pI8&_nc_zt=23&_nc_ht=scontent.fbkk13-3.fna&_nc_gid=GbljdPLRgUnHjSy5hUvxuQ&oh=00_AYFuCWoRPpzxy2pVmWRf1Jre8w_oZcyfo4P1O_9laZk1Ow&oe=67F02F1E', 1, 1),
    (1, 'Physics', 'Physics Lab under Science Faculty', 'Physics Lab', 'https://science.ku.ac.th/physics', 'https://physics.sci.ku.ac.th/sites/default/files/logo_0.png', 1, 1),
    (1, 'Chemistry', 'Chemistry Lab under Science Faculty', 'Chemistry Lab', 'https://science.ku.ac.th/chemistry', 'https://scontent.fbkk13-1.fna.fbcdn.net/v/t39.30808-6/223551850_4302772183122236_2252554062280833551_n.jpg?_nc_cat=105&ccb=1-7&_nc_sid=6ee11a&_nc_ohc=PzcUOYgXQt8Q7kNvgHTfDD0&_nc_oc=Admaz9DZJZr-wjMpJKvVmk0sd-sBj9U70E32zuUGzSnULVXWcEpYSoVdnCcdYL_QqmjEaTekpMGxIODo-KWR5s8u&_nc_zt=23&_nc_ht=scontent.fbkk13-1.fna&_nc_gid=ge6XkWrZjLYpGHKL-nGu5w&oh=00_AYFlMI71_Ft6i1XB5lxezXtE3bwfS25kJWgeLY9Qqzt-3g&oe=67F01811', 1, 1),
    (1, 'Biology', 'Biology Research Center under Science Faculty', 'Biology Research', 'https://science.ku.ac.th/biology', 'https://scontent.fbkk9-2.fna.fbcdn.net/v/t39.30808-6/278839740_109633338376594_3899151675007542408_n.jpg?_nc_cat=109&ccb=1-7&_nc_sid=6ee11a&_nc_ohc=Om1BeXzhFSIQ7kNvgHPIary&_nc_oc=Admta2jhepGEo7UaXqLkYJ0NYM2LPuYg9KRv0KT9pv3Gt3tywazS0pIBwcHdj1M6BKBlR_zBup49sdgWeljLfhks&_nc_zt=23&_nc_ht=scontent.fbkk9-2.fna&_nc_gid=410mzNziX4XjqQ3gRIQ7vw&oh=00_AYElDSbgueM08KYRJFbMUESUIYb_3TW89GOSf71YYDGgBA&oe=67F00AFC', 1, 1),
    (1, 'Office', 'Office for Computer Science Department', 'CS Office', 'https://science.ku.ac.th/cs-office', 'https://scontent.fbkk8-2.fna.fbcdn.net/v/t39.30808-6/300799525_508793294582923_4415974765950681482_n.jpg?_nc_cat=110&ccb=1-7&_nc_sid=6ee11a&_nc_ohc=Xm1M91mkxYUQ7kNvgGWaMYf&_nc_oc=Adk3a8mAfu2eqbi9NcvRj8e-eVF6cobuhXXTmFo-_XIlH6T3SSK_dMtz_VeB1C8mCsq4Bg1B65NGruoYvYuiUN57&_nc_zt=23&_nc_ht=scontent.fbkk8-2.fna&_nc_gid=5joOd500UlUwKv-MeFlsSQ&oh=00_AYHLH8280tOgsJS9GU-hECGbe27tJn4V0V_V2fzN0X0mCg&oe=67EFFC3C', 1, 1),
    (1, 'Nisit', 'Student workspace under Computer Science', 'CS Nisit', 'https://science.ku.ac.th/cs-nisit', 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQsm78Pgbz4DidfAOCYSMmAnsXC8oSLen8C5A&s', 1, 1);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sites;
-- +goose StatementEnd
