package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/romitou/gostats/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strings"
)

type GuildData struct {
	Name          string `json:"name"`
	MemberCount   int    `json:"member_count"`
	PresenceCount int    `json:"presence_count"`
}

// Fetch member_count and presence_count from Discord's API for guildId
func GetGuildData(guildId string) GuildData {
	req, err := http.NewRequest("GET", "https://discord.com/api/v9/guilds/"+guildId+"/preview", nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("DISCORD_BOT_TOKEN"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	var guildData GuildData
	err = json.NewDecoder(resp.Body).Decode(&guildData)
	if err != nil {
		log.Fatalln(err)
	}
	return guildData
}

func FormattedName(name string) string {
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, ":", "")
	name = strings.ReplaceAll(name, "'", "")
	name = strings.ReplaceAll(name, "!", "")
	name = strings.ReplaceAll(name, "-", "")
	name = strings.ReplaceAll(name, "&", "")
	return strings.ToLower(name)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_DSN")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	guildIds := strings.Split(os.Getenv("DISCORD_GUILD_IDS"), ",")
	for _, guildId := range guildIds {
		guildData := GetGuildData(guildId)
		db.Table("guild_" + FormattedName(guildData.Name)).Create(&models.DiscordStat{
			MemberCount:   guildData.MemberCount,
			PresenceCount: guildData.PresenceCount,
		})
	}
}
