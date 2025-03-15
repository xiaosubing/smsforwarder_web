package service

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func GetMessages(c *gin.Context) {
	// 用于存储随机获取的消息
	var messages []map[string]string
	// 获取 5 次随机消息
	for i := 0; i < 5; i++ {
		sender, content := randData()
		// 将每次获取的消息封装成 map 并添加到 messages 切片中
		message := map[string]string{
			"sender":  sender,
			"content": content,
		}
		messages = append(messages, message)
	}
	// 以 JSON 格式返回消息列表
	c.JSON(http.StatusOK, messages)
}

func GetPhones(c *gin.Context) {
	var phones []string
	for i := 0; i < 5; i++ {
		p := randPhone()
		phones = append(phones, p)
	}
	c.JSON(http.StatusOK, phones)
}

func randData() (string, string) {
	messages := map[string]string{
		"10010":     "【联通云盘】您通过10010得的联通云盘畅享会员月卡奖品已到账，您可点击 https://pan.wo.cn/renew/su/99LRpBB1qVY 前往联通云盘APP使用。",
		"10086":     "【移动云盘】您通过转存有10086 盘APP使用。",
		"10000":     "【电信云盘】您通过1000000 奖品",
		"1002102":   "【联通云盘】您通过10010得的联通云盘畅享会员月卡奖品已到账，您可hahahhLRpBB1qVY 前往联通云盘APP使用。",
		"100862":    "【移动云盘】您通过转存有10086 盘APP使用。",
		"1003":      "【电信云盘】您通过1000000 奖品",
		"10040":     "【【飞鱼CRM】您好，杭州千尊信息技术有限公司公司客服回访专员给您回访来电，请放心接听",
		"104r246":   "【移动云盘】您通过转存有10086 盘APP使用。",
		"100t4":     "【电信云盘】您通过1000 奖品盘畅享会员 奖品",
		"1004t4":    "【联通云盘】您通过10010得的联通云盘畅享会员通云盘畅享会员月卡奖品已到账，您可点击 https://pan.wo.cn/renew/su/99LRpBB1qVY 前往联通云盘APP使用。",
		"100t42":    "【移动云盘】您通过转存有10086 盘APP使用。",
		"104t23t3":  "【【积3r3r3tr2t3至本月08日，您t23tt23t23t23t23t224积分。",
		"100r23r23": "【联通云盘】您通345t4https://pan.wo.cnt4tt2云盘APP使用。",
		"100332":    "【移动云盘】您2e盘APP使用。",
		"10tr342":   "【【积分提醒】截至本月08日，您有824积分。",
		"10t233t2":  "【联通云盘】您验证码：631140，短信验证码2分钟内有效，请勿泄漏给他人。【中国联通】",
		"100t23":    "【移动云盘】您通过转存有10086 盘APP使用。",
		"10tfq2323": "【电信云盘】您通过1000000 奖品",
	}

	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	// 存储所有键的切片
	keys := make([]string, 0, len(messages))
	for key := range messages {
		keys = append(keys, key)
	}

	// 生成随机索引
	randomIndex := rand.Intn(len(keys))

	// 获取随机键
	randomKey := keys[randomIndex]
	return randomKey, messages[randomKey]
}

func randPhone() string {
	messages := []string{
		"13012345678", "13012345679",
		"13012345677", "13012345676", "13012345675", "13012345674", "13012345678",
	}
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	// 生成随机索引
	randomIndex := rand.Intn(len(messages))

	// 获取随机键
	randomKey := messages[randomIndex]
	return randomKey
}
