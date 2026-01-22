package pdb

import (
	"os"
	"path"
	"pdb_proxy/conf"

	"github.com/gofiber/fiber/v2"
)

func PdbQuery(c *fiber.Ctx) error {
	pdbName := c.Params("pdbname")
	pdbHash := c.Params("pdbhash")
	pdbQuery := pdbName + "/" + pdbHash + "/" + pdbName
	pdbPath := path.Join(conf.PdbDir, pdbQuery)
	//log.Printf("Pdb Path: %s", pdbPath)

	_, err := os.Stat(pdbPath)
	if err == nil {
		return c.SendFile(pdbPath)
	} else {
		pdbUrl := conf.PdbServer + "/" + pdbQuery
		err := DownLoadFile(pdbUrl, pdbPath)
		if err != nil {
			os.Remove(pdbPath)
			return c.Status(404).SendString("")
		} else {
			return c.SendFile(pdbPath)
		}
	}

}
