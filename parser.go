package main

import (
	"sort"

	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Fortress.2021.BluRay.1080p.AVC.DTS-HD.MA5.1-MTeam
type MediaInfo struct {
	Title      string   `json:"title"`       // 从视频提取的完整文件名 鹰眼 Hawkeye
	Year       int      `json:"year"`        // 年份：2020、2021
	Source     string   `json:"source"`      // 来源：WEB-DL、BluRay、HDTV、DVD、BD、UHD BluRay
	Resolution string   `json:"resolution"`  // 分辨率：4K、1080P、720P、SD
	Studio     string   `json:"studio"`      // 制作公司：MTeam、CHD、HDChina
	Channel    string   `json:"channel"`     // 渠道：Netflix、Disney+
	Version    string   `json:"version"`     // 版本：Remastered、Extended Edition
	Season     string   `json:"season"`      // 季：S01、S02
	Episode    string   `json:"episode"`     // 集：E01、E02
	Type       string   `json:"type"`        // 类型：movie、tv、anime
	AudioCodec string   `json:"audio_codec"` // 音频编码：DTS、DTS-HD、DDP5.1、AAC、AAC(5.1)
	VideoCodec string   `json:"video_codec"` // 视频编码：x264、x265、H.265、VP9
	Languages  []string `json:"languages"`   // 语言：chs、eng、japanese
}

var (
	video = []string{
		"mkv",
		"mp4",
		"ts",
		"avi",
		"wmv",
		"m4v",
		"flv",
		"webm",
		"mpeg",
		"mpg",
		"3gp",
		"3gpp",
		"ts",
		"iso",
		"mov",
	}
	audioCodecs = []string{
		"AAC",
		"AAC(5.1)",
		"DTS",
		"DDP5.1",
		"AC3.4Audio",
		"DTS.2Audios",
		"DTS-HD.MA5.1",
		"DTS-HD",
	}
	videoCodecs = []string{
		"AVC",
		"X264",
		"X265",
		"Webm",
		"VP9",
		"H.265",
	}
	source = []string{
		"WEB-DL",
		"Blu-Ray",
		"BluRay",
		"HDTV",
		"CCTVHD",
	}
	studio = []string{
		"hmax",
		"netflix",
		"funimation",
		"amzn",
		"hulu",
		"kktv",
		"crunchyroll",
		"bbc",
		"greenotea",
		"nowys",
		"hdchina",
		"cmct",
		"hdwing",
	}
	versions = []string{
		"remastered",
		"extended.edition",
	}
	languages = []string{
		"chs",
		"eng",
		"japanese",
	}
	delimiter = []string{
		"-",
		".",
		",",
		"_",
		" ",
		"[",
		"]",
		"(",
		")",
		"{",
		"}",
		"@",
		":",
		"：",
	}
	delimiterExecute = []string{
		"AAC(5.1)",
		"AC3.4Audio",
		"DTS.2Audios",
		"DTS-HD.MA5.1",
		"WEB-DL",
		"DDP 5.1",
		"DDP5.1",
		"DDP.5.1",
		"H.265",
		"Blu-Ray",
		"MA5.1",
		"MA 5.1",
		"MA.5.1",
		"MA7.1",
		"MA 7.1",
		"MA.7.1",
		"DTS-HD",
		"Extended.Edition",
	}
	channel = []string{
		"OAD",
		"OVA",
		"BD",
		"DVD",
		"SP",
	}

	collectionMatch, _    = regexp.Compile("[sS](0|)[0-9]+-[sS](0|)[0-9]+")
	yearRangeLikeMatch, _ = regexp.Compile("[12][0-9]{3}-[12][0-9]{3}")
	yearRangeMatch, _     = regexp.Compile("[12][0-9]{3}-[12][0-9]{3}")
	yearMatch, _          = regexp.Compile("^[12][0-9]{3}$")
	resolutionMatch, _    = regexp.Compile("[0-9]{3,4}Xx*[0-9]{3,4}")
	resolutionMatch2, _   = regexp.Compile("([0-9]+[pPiI]|[24][kK])")
	episodeMatch, _       = regexp.Compile(`(?i)((?:s|第|season|s|S)([0-9]+)(?:季|))?((?:第|episode|e|ep|E)([0-9]+)(?:集|))$`)
)

func SplitWith(s string, sep []string, exclude []string) []string {
	// 对 exclude 中的关键字按长度排序
	sort.Slice(exclude, func(i, j int) bool {
		return len(exclude[i]) > len(exclude[j])
	})

	// 创建映射，将 exclude 中的关键字替换为特殊标记
	replaceMap := make(map[string]string)
	for i, word := range exclude {
		replaceMap[word] = fmt.Sprintf("<%d>", i+1)
	}

	// 替换需要排除的字符串为特殊标记
	for word, replacement := range replaceMap {
		// 使用正则的 i 选项进行大小写不敏感的替换
		re := regexp.MustCompile("(?i)" + regexp.QuoteMeta(word))
		s = re.ReplaceAllString(s, replacement)
	}

	// 使用分隔符进行分割
	splits := strings.FieldsFunc(s, func(r rune) bool {
		for _, separator := range sep {
			if strings.ContainsRune(separator, r) {
				return true
			}
		}
		return false
	})

	// 还原特殊标记为原来的字符串
	for word, replacement := range replaceMap {
		for j, part := range splits {
			splits[j] = strings.ReplaceAll(part, replacement, word)
		}
	}

	return splits
}

// IsCollection 是否是合集，如S01-S03季
func IsCollection(name string) bool {
	return collectionMatch.MatchString(name) || yearRangeMatch.MatchString(name)
}

// IsYearRangeLike 判断并返回年范围，用于合集
func IsYearRangeLike(name string) string {
	return yearRangeLikeMatch.FindString(name)
}

// IsYearRange 判断并返回年范围，用于合集
func IsYearRange(name string) string {
	return yearRangeMatch.FindString(name)
}

// IsYear 判断是否是年份
func IsYear(name string) int {
	if !yearMatch.MatchString(name) {
		return 0
	}
	year, _ := strconv.Atoi(name)
	return year
}

// IsSource 片源
func IsSource(name string) string {
	for _, item := range source {
		if strings.EqualFold(item, name) {
			return name
		}
	}
	return ""
}

// IsStudio 发行公司
func IsStudio(name string) string {
	for _, item := range studio {
		if strings.EqualFold(item, name) {
			return name
		}
	}
	return ""
}

// IsChannel 发行渠道
func IsChannel(name string) string {
	for _, item := range channel {
		if strings.EqualFold(item, name) {
			return name
		}
	}
	return ""
}

func IsVersion(name string) string {
	for _, item := range versions {
		if strings.EqualFold(item, name) {
			return name
		}
	}
	return ""
}

func IsLanguage(name string) string {
	for _, item := range languages {
		if strings.EqualFold(item, name) {
			return name
		}
	}
	return ""
}

func IsAudioCodec(name string) string {
	for _, item := range audioCodecs {
		if strings.EqualFold(item, name) {
			return name
		}
	}
	return ""
}

func IsVideoCodec(name string) string {
	for _, item := range videoCodecs {
		if strings.EqualFold(item, name) {
			return name
		}
	}
	return ""
}

// IsResolution 分辨率
func IsResolution(name string) string {
	if r := resolutionMatch2.FindString(name); r != "" {
		return r
	}
	return resolutionMatch.FindString(name)
}

func IsType(name string) string {
	for _, item := range video {
		if strings.EqualFold(item, name) {
			return item
		}
	}
	return ""
}

// Split 影视目录或文件名切割
func Split(name string) []string {
	return SplitWith(name, delimiter, delimiterExecute)
}

// MatchEpisode 匹配季和集
func MatchEpisode(name string) (ok bool, season, episode string) {
	find := episodeMatch.FindStringSubmatch(name)
	ok = len(find) == 5
	if !ok {
		return
	}
	season, episode = find[2], find[4]
	season = fmt.Sprintf("S%02s", season)
	episode = fmt.Sprintf("E%02s", episode)
	return
}

func Parse(filename string) (info MediaInfo) {
	nameStart := false
	nameStop := false
	parts := Split(filename)
	for _, item := range parts {
		if year := IsYear(item); year > 0 {
			info.Year = year
			nameStop = true
			continue
		}
		if codec := IsAudioCodec(item); len(codec) > 0 {
			info.AudioCodec = codec
			nameStop = true
		}

		if codec := IsVideoCodec(item); len(codec) > 0 {
			info.VideoCodec = codec
			nameStop = true
		}

		if source := IsSource(item); len(source) > 0 {
			info.Source = source
			nameStop = true
			continue
		}

		if studio := IsStudio(item); len(studio) > 0 {
			info.Studio = studio
			nameStop = true
			continue
		}

		if channel := IsChannel(item); len(channel) > 0 {
			info.Channel = channel
			nameStop = true
			continue
		}

		if version := IsVersion(item); len(version) > 0 {
			info.Version = version
			nameStop = true
		}

		if language := IsLanguage(item); len(language) > 0 {
			info.Languages = append(info.Languages, language)
			nameStop = true
		}

		if resolution := IsResolution(item); len(resolution) > 0 {
			info.Resolution = resolution
			nameStop = true
			continue
		}

		if typ := IsType(item); len(typ) > 0 {
			info.Type = typ
			nameStop = true
			continue
		}

		ok, s, e := MatchEpisode(item)
		if ok {
			info.Season = s
			info.Episode = e
			nameStop = true
			continue
		}

		if !nameStart {
			nameStart = true
			nameStop = false
		}

		if !nameStop {
			info.Title += item + " "
		}
	}
	info.Title = strings.TrimSpace(info.Title)
	return
}
