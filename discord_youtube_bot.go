package main
import (
  "fmt"
  "os"
  "os/signal"
  "syscall"
  "strings"
 
  "github.com/bwmarrin/discordgo"
)
 
const(
  TOKEN = "NzQ5MDcxMjI3MzA5NTIyOTc2.X0mo8Q.f5XfW7f2_hWNQO8Nzv4eu8HuJOQ"
)
 
func main() {
  dg, err := discordgo.New("Bot " + TOKEN)
  if err != nil {
    fmt.Println("error:start\n", err)
    return
  }
 
  //on message
  dg.AddHandler(messageCreate)
 
  err = dg.Open()
  if err != nil {
    fmt.Println("error:wss\n", err)
    return
  }
  fmt.Println("BOT Running...")
 
  //シグナル受け取り可にしてチャネル受け取りを待つ（受け取ったら終了）
  sc := make(chan os.Signal, 1)
  signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
  <-sc
  dg.Close()
}
 
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
  if m.Author.Bot {
    return
  }
  nick := m.Author.Username
  member, err := s.State.Member(m.GuildID, m.Author.ID)
  if err == nil && member.Nick != "" {
    nick = member.Nick
  }
  fmt.Println("< "+m.Content+" by "+nick)
 
  if m.Content == "おはよう" {
    s.ChannelMessageSend(m.ChannelID, "おはよう朝日です！ただいま9時20分")
    fmt.Println(">おはよう朝日です！ただいま9時20分")
  }
  if strings.Contains(m.Content,"美人") {
    s.ChannelMessageSend(m.ChannelID, "美人動画内見？https://www.youtube.com/watch?v=EzmYhkJE638")
    fmt.Println("> 美人動画内見")
  }
}