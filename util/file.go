package util

import (
	"fmt"
	"os"
)

func FileOrDirExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func FileOrDirIsNotExists(path string) bool {
	return !FileOrDirExists(path)
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	return !IsDir(path)
}

func MkDirForce(path string) error {
	if FileOrDirIsNotExists(path) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println("创建文件夹失败,error info:", err)
			return err
		}
		return err
	}
	return nil
}

func GetFileExtByHttpContentType(contentType string) string {
	ext := ""
	switch contentType {
	case "application/epub+zip":
		ext = ".epub"
	case "application/fractals":
		ext = ".fif"
	case "application/futuresplash":
		ext = ".spl"
	case "application/hta":
		ext = ".hta"
	case "application/mac-binhex40":
		ext = ".hqx"
	case "application/ms-vsi":
		ext = ".vsi"
	case "application/msaccess":
		ext = ".accdb"
	case "application/msaccess.addin":
		ext = ".accda"
	case "application/msaccess.cab":
		ext = ".accdc"
	case "application/msaccess.exec":
		ext = ".accde"
	case "application/msaccess.ftemplate":
		ext = ".accft"
	case "application/msaccess.runtime":
		ext = ".accdr"
	case "application/msaccess.template":
		ext = ".accdt"
	case "application/msaccess.webapplication":
		ext = ".accdw"
	case "application/msonenote":
		ext = ".one"
	case "application/msword":
		ext = ".doc"
	case "application/opensearchdescription+xml":
		ext = ".osdx"
	case "application/pdf":
		ext = ".pdf"
	case "application/pkcs10":
		ext = ".p10"
	case "application/pkcs7-mime":
		ext = ".p7c"
	case "application/pkcs7-signature":
		ext = ".p7s"
	case "application/pkix-cert":
		ext = ".cer"
	case "application/pkix-crl":
		ext = ".crl"
	case "application/postscript":
		ext = ".ps"
	case "application/vnd.ms-excel":
		ext = ".xls"
	case "application/vnd.ms-excel.12":
		ext = ".xlsx"
	case "application/vnd.ms-excel.addin.macroEnabled.12":
		ext = ".xlam"
	case "application/vnd.ms-excel.sheet.binary.macroEnabled.12":
		ext = ".xlsb"
	case "application/vnd.ms-excel.sheet.macroEnabled.12":
		ext = ".xlsm"
	case "application/vnd.ms-excel.template.macroEnabled.12":
		ext = ".xltm"
	case "application/vnd.ms-officetheme":
		ext = ".thmx"
	case "application/vnd.ms-pki.certstore":
		ext = ".sst"
	case "application/vnd.ms-pki.pko":
		ext = ".pko"
	case "application/vnd.ms-pki.seccat":
		ext = ".cat"
	case "application/vnd.ms-powerpoint":
		ext = ".ppt"
	case "application/vnd.ms-powerpoint.12":
		ext = ".pptx"
	case "application/vnd.ms-powerpoint.addin.macroEnabled.12":
		ext = ".ppam"
	case "application/vnd.ms-powerpoint.presentation.macroEnabled.12":
		ext = ".pptm"
	case "application/vnd.ms-powerpoint.slide.macroEnabled.12":
		ext = ".sldm"
	case "application/vnd.ms-powerpoint.slideshow.macroEnabled.12":
		ext = ".ppsm"
	case "application/vnd.ms-powerpoint.template.macroEnabled.12":
		ext = ".potm"
	case "application/vnd.ms-publisher":
		ext = ".pub"
	case "application/vnd.ms-visio.viewer":
		ext = ".vsd"
	case "application/vnd.ms-word.document.12":
		ext = ".docx"
	case "application/vnd.ms-word.document.macroEnabled.12":
		ext = ".docm"
	case "application/vnd.ms-word.template.12":
		ext = ".dotx"
	case "application/vnd.ms-word.template.macroEnabled.12":
		ext = ".dotm"
	case "application/vnd.ms-wpl":
		ext = ".wpl"
	case "application/vnd.ms-xpsdocument":
		ext = ".xps"
	case "application/vnd.oasis.opendocument.presentation":
		ext = ".odp"
	case "application/vnd.oasis.opendocument.spreadsheet":
		ext = ".ods"
	case "application/vnd.oasis.opendocument.text":
		ext = ".odt"
	case "application/vnd.openxmlformats-officedocument.presentationml.presentation":
		ext = ".pptx"
	case "application/vnd.openxmlformats-officedocument.presentationml.slide":
		ext = ".sldx"
	case "application/vnd.openxmlformats-officedocument.presentationml.slideshow":
		ext = ".ppsx"
	case "application/vnd.openxmlformats-officedocument.presentationml.template":
		ext = ".potx"
	case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":
		ext = ".xlsx"
	case "application/vnd.openxmlformats-officedocument.spreadsheetml.template":
		ext = ".xltx"
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
		ext = ".docx"
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.template":
		ext = ".dotx"
	case "application/windows-appcontent+xml":
		ext = ".appcontent-ms"
	case "application/x-compress":
		ext = ".z"
	case "application/x-compressed":
		ext = ".solitairetheme8"
	case "application/x-dtcp1":
		ext = ".dtcp-ip"
	case "application/x-gzip":
		ext = ".gz"
	case "application/x-itunes-itls":
		ext = ".itls"
	case "application/x-itunes-itms":
		ext = ".itms"
	case "application/x-itunes-itpc":
		ext = ".itpc"
	case "application/x-jtx+xps":
		ext = ".jtx"
	case "application/x-latex":
		ext = ".latex"
	case "application/x-mix-transfer":
		ext = ".nix"
	case "application/x-mplayer2":
		ext = ".asx"
	case "application/x-ms-application":
		ext = ".application"
	case "application/x-ms-vsto":
		ext = ".vsto"
	case "application/x-ms-wmd":
		ext = ".wmd"
	case "application/x-ms-wmz":
		ext = ".wmz"
	case "application/x-ms-xbap":
		ext = ".xbap"
	case "application/x-mswebsite":
		ext = ".website"
	case "application/x-pkcs12":
		ext = ".p12"
	case "application/x-pkcs7-certificates":
		ext = ".p7b"
	case "application/x-pkcs7-certreqresp":
		ext = ".p7r"
	case "application/x-podcast":
		ext = ".pcast"
	case "application/x-shockwave-flash":
		ext = ".swf"
	case "application/x-stuffit":
		ext = ".sit"
	case "application/x-tar":
		ext = ".tar"
	case "application/x-troff-man":
		ext = ".man"
	case "application/x-wmplayer":
		ext = ".asx"
	case "application/x-x509-ca-cert":
		ext = ".cer"
	case "application/x-zip-compressed":
		ext = ".zip"
	case "application/xaml+xml":
		ext = ".xaml"
	case "application/xhtml+xml":
		ext = ".xht"
	case "application/xml":
		ext = ".xml"
	case "application/zip":
		ext = ".zip"
	case "audio/3gpp":
		ext = ".3gp"
	case "audio/3gpp2":
		ext = ".3g2"
	case "audio/aac":
		ext = ".aac"
	case "audio/aiff":
		ext = ".aiff"
	case "audio/amr":
		ext = ".amr"
	case "audio/basic":
		ext = ".au"
	case "audio/ec3":
		ext = ".ec3"
	case "audio/l16":
		ext = ".lpcm"
	case "audio/mid":
		ext = ".mid"
	case "audio/midi":
		ext = ".mid"
	case "audio/mp3":
		ext = ".mp3"
	case "audio/mp4":
		ext = ".m4a"
	case "audio/MP4A-LATM":
		ext = ".m4a"
	case "audio/mpeg":
		ext = ".mp3"
	case "audio/mpegurl":
		ext = ".m3u"
	case "audio/mpg":
		ext = ".mp3"
	case "audio/vnd.dlna.adts":
		ext = ".adts"
	case "audio/vnd.dolby.dd-raw":
		ext = ".ac3"
	case "audio/wav":
		ext = ".wav"
	case "audio/x-aiff":
		ext = ".aiff"
	case "audio/x-flac":
		ext = ".flac"
	case "audio/x-m4a":
		ext = ".m4a"
	case "audio/x-m4r":
		ext = ".m4r"
	case "audio/x-matroska":
		ext = ".mka"
	case "audio/x-mid":
		ext = ".mid"
	case "audio/x-midi":
		ext = ".mid"
	case "audio/x-mp3":
		ext = ".mp3"
	case "audio/x-mpeg":
		ext = ".mp3"
	case "audio/x-mpegurl":
		ext = ".m3u"
	case "audio/x-mpg":
		ext = ".mp3"
	case "audio/x-ms-wax":
		ext = ".wax"
	case "audio/x-ms-wma":
		ext = ".wma"
	case "audio/x-wav":
		ext = ".wav"
	case "image/bmp":
		ext = ".dib"
	case "image/gif":
		ext = ".gif"
	case "image/jpeg":
		ext = ".jpg"
	case "image/jps":
		ext = ".jps"
	case "image/mpo":
		ext = ".mpo"
	case "image/pjpeg":
		ext = ".jpg"
	case "image/png":
		ext = ".png"
	case "image/pns":
		ext = ".pns"
	case "image/svg+xml":
		ext = ".svg"
	case "image/tiff":
		ext = ".tif"
	case "image/vnd.ms-dds":
		ext = ".dds"
	case "image/vnd.ms-photo":
		ext = ".wdp"
	case "image/x-emf":
		ext = ".emf"
	case "image/x-icon":
		ext = ".ico"
	case "image/x-png":
		ext = ".png"
	case "image/x-wmf":
		ext = ".wmf"
	case "midi/mid":
		ext = ".mid"
	case "model/vnd.dwfx+xps":
		ext = ".dwfx"
	case "model/vnd.easmx+xps":
		ext = ".easmx"
	case "model/vnd.edrwx+xps":
		ext = ".edrwx"
	case "model/vnd.eprtx+xps":
		ext = ".eprtx"
	case "pkcs10":
		ext = ".p10"
	case "pkcs7-mime":
		ext = ".p7m"
	case "pkcs7-signature":
		ext = ".p7s"
	case "pkix-cert":
		ext = ".cer"
	case "pkix-crl":
		ext = ".crl"
	case "text/calendar":
		ext = ".ics"
	case "text/css":
		ext = ".css"
	case "text/directory":
		ext = ".vcf"
	case "text/directory;profile=vCard":
		ext = ".vcf"
	case "text/html":
		ext = ".html"
	case "text/plain":
		ext = ".txt"
	case "text/scriptlet":
		ext = ".wsc"
	case "text/vcard":
		ext = ".vcf"
	case "text/x-component":
		ext = ".htc"
	case "text/x-ms-contact":
		ext = ".contact"
	case "text/x-ms-iqy":
		ext = ".iqy"
	case "text/x-ms-odc":
		ext = ".odc"
	case "text/x-ms-rqy":
		ext = ".rqy"
	case "text/x-vcard":
		ext = ".vcf"
	case "text/xml":
		ext = ".xml"
	case "video/3gpp":
		ext = ".3gpp"
	case "video/3gpp2":
		ext = ".3gp2"
	case "video/avi":
		ext = ".avi"
	case "video/mp4":
		ext = ".mp4"
	case "video/mpeg":
		ext = ".mpeg"
	case "video/mpg":
		ext = ".mpeg"
	case "video/msvideo":
		ext = ".avi"
	case "video/quicktime":
		ext = ".mov"
	case "video/vnd.dece.mp4":
		ext = ".uvu"
	case "video/vnd.dlna.mpeg-tts":
		ext = ".tts"
	case "video/wtv":
		ext = ".wtv"
	case "video/x-m4v":
		ext = ".m4v"
	case "video/x-matroska":
		ext = ".mkv"
	case "video/x-mpeg":
		ext = ".mpeg"
	case "video/x-mpeg2a":
		ext = ".mpeg"
	case "video/x-ms-asf":
		ext = ".asx"
	case "video/x-ms-asf-plugin":
		ext = ".asx"
	case "video/x-ms-dvr":
		ext = ".dvr-ms"
	case "video/x-ms-wm":
		ext = ".wm"
	case "video/x-ms-wmv":
		ext = ".wmv"
	case "video/x-ms-wmx":
		ext = ".wmx"
	case "video/x-ms-wvx":
		ext = ".wvx"
	case "video/x-msvideo":
		ext = ".avi"
	case "vnd.ms-pki.certstore":
		ext = ".sst"
	case "vnd.ms-pki.pko":
		ext = ".pko"
	case "vnd.ms-pki.seccat":
		ext = ".cat"
	case "x-pkcs12":
		ext = ".p12"
	case "x-pkcs7-certificates":
		ext = ".p7b"
	case "x-pkcs7-certreqresp":
		ext = ".p7r"
	case "application/vnd.android.package-archive":
		ext = ".apk"
	case "application/vnd.android.obb":
		ext = ".obb"
	case "x-x509-ca-cert":
		ext = ".cer"
	case "application/json":
		ext = ".json"
	}
	return ext
}
