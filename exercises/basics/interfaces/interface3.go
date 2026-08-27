	package main
	import "fmt"

	type JSONExporter struct{
		FileName string
	}

	type CSVExporter struct{
		FileName string
	}

	type Exporter interface{
		GetFormat() string
		Export(data string) string

	}
	func main(){

		fileName:=JSONExporter{
			FileName: "data.json",
		}
		fileName2:=CSVExporter{
			FileName: "data.csv",
		}

		ProccessExport(fileName,"Öğrenci notları")
		ProccessExport(fileName2,"Öğrenci notları2")
	}

	func (j JSONExporter) GetFormat()string{
		return "JSON"
	}

	func (j JSONExporter) Export(data string) string{
		return fmt.Sprintf("%s verisi %s dosyasına aktarıldı",data,j.FileName)
	}

	func (c CSVExporter) GetFormat() string{
		return "CSV"
	}

	func (c CSVExporter) Export(data string) string{
		return fmt.Sprintf("%s verisi %s dosyasına aktarıldı",data,c.FileName)
	}

	func ProccessExport(e Exporter,data string){
		fmt.Println("Export format:",e.GetFormat())
		fmt.Println(e.Export(data))
	}