package models

import (
	"MeNA-Api/common"
	"MeNA-Api/models/database"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

func AddMusicTag(content string) (tag_id int64, err error) {
	var db_orm = orm.NewOrm() //orm不能定义为全局的
	music_tag := &db.AllMusicTags{
		TagName: content,
	}
	err = db_orm.Read(music_tag, "TagName")
	if err != nil || music_tag.TagName == "" {
		music_tag = &db.AllMusicTags{
			TagName: content,
			Created: time.Now(),
			Update:  time.Now(),
		}
		v, e := db_orm.Insert(music_tag)
		if e != nil {
			common.WriteLog("add new music_tag error @MusicMod.AddMusicTag() : " + err.Error())
			return 0, e
		}
		return v, nil
	}
	return -1, nil // 已经存在该tag不能添加重复的tag

}

func GetOneMusicTag(tag_id int64) (music_tag db.AllMusicTags, err error) {
	var db_orm = orm.NewOrm()
	music_tag = db.AllMusicTags{
		Id: tag_id,
	}
	err = db_orm.Read(&music_tag)
	if err != nil {
		// common.WriteLog("get one character error @characterMod.GetOneCharacter() : " + e.Error())
		return music_tag, errors.New("Music-Tag-Id Not Exist")
	}
	return music_tag, nil
}

func GetAllMusicTag(opts map[string]int) (Result, error) {
	var data Result
	// var data_map = make(map[string]string)
	o := orm.NewOrm()
	// rows := make([]db.AllMusicTags, 0)
	var rows []db.AllMusicTags
	var list = make([]string, opts["limit"])
	/************放弃使用原生sql查询************/
	// sql := "SELECT tag_name FROM all_music_tags limit " + opts["limit"] + " offset " + opts["offset"]
	// num, err := o.Raw(sql).Values(&obj)
	/******************************************/
	qs := o.QueryTable("all_music_tags")
	num, err := qs.All(&rows, "TagName")
	if err == nil {
		tmpint, _ := strconv.Atoi(strconv.FormatInt(num, 10))
		check := (tmpint % opts["limit"])
		var total_pages int
		if check == 0 {
			total_pages = tmpint / opts["limit"]
		} else {
			total_pages = (tmpint / opts["limit"]) + 1
		}
		data.TotalCount = tmpint
		data.TotalPages = total_pages
		// data_map["total_count"] = strconv.Itoa(tmpint)
		// data_map["total_pages"] = strconv.Itoa(total_pages)
		// fmt.Println("===total_pages===")
		// fmt.Println(total_pages)
	}
	num, err = qs.Limit(opts["limit"], opts["offset"]).All(&rows, "TagName")
	if err == nil && num > 0 {
		// fmt.Println("===rows===")
		// fmt.Println(rows)
		// fmt.Println(num)
		i := 0
		for _, m := range rows {
			// fmt.Println("===in for range===")
			// fmt.Println(reflect.TypeOf(m.TagName))
			// fmt.Println(m.TagName)
			list[i] = m.TagName
			i++
		}
		data.List = list

		//todo go 语言接口
		// jsonbyte, _ := json.Marshal(list)
		// list_str := string(jsonbyte)
		// data_map["list"] = list_str
	} else {
		common.WriteLog("select all music_tag error @MusicMod.GetAllMusicTag() : " + err.Error())
	}
	fmt.Println("===data===")
	fmt.Println(data)
	return data, err
	// fmt.Println("===data===")
	// fmt.Println(data)
}

// func GetAllCharacter() map[string]*db.Character {
// 	return Characters
// }

// func UpdateCharacter(character_id string, content string) (err error) {
// 	if v, ok := Characters[CharacterId]; ok {
// 		v.Score = Score
// 		return nil
// 	}
// 	return errors.New("CharacterId Not Exist")
// }

// func DeleteCharacter(character_id string) {
// 	delete(Characters, CharacterId)
// }
