package scrape

type ScrapeCfg struct {
	Username string
	Password string
	Url      string
	DestDir  string
}

func Run(cfg *ScrapeCfg) error {
	return nil
}

func fetchPage() {

}
