package common

import (
	"os"
	"strings"
	"testing"
)

// GetAppConfig test
func TestGetAppConfig(t *testing.T) {
	appConfig := AppConfig{LogLevel: "$LOG_LEVEL"}
	ymlConfig := YmlConfig{App: appConfig}

	// Test data
	os.Setenv("LOG_LEVEL", "test")

	if !strings.EqualFold(ymlConfig.GetAppConfig().LogLevel, "test") {
		t.Errorf("Incorrect GetAppConfig test. log_level = %s", ymlConfig.GetAppConfig().LogLevel)
		t.FailNow()
	}
}

// GetCacherConfig test
func TestGetCacherConfig(t *testing.T) {
	cacherConfig := CacherConfig{TimeMillisStr: "$CACHER_TIME_MILLIS"}
	ymlConfig := YmlConfig{Cacher: cacherConfig}

	// Test data
	os.Setenv("CACHER_TIME_MILLIS", "100")

	if !strings.EqualFold(ymlConfig.GetCacherConfig().TimeMillisStr, "100") {
		t.Errorf("Incorrect CacherConfig test. time-millis = %s", ymlConfig.GetCacherConfig().TimeMillisStr)
		t.FailNow()
	}
}

// GetServerConfig test
func TestGetServerConfig(t *testing.T) {
	serverConfig := ServerConfig{
		Port:                     "$SERVER_PORT",
		SignedInPrivateKeyBase64: "$SERVER_PRIVATE_KEY",
		TokenExpireHourStr:       "$SERVER_TOKEN_EXPIRE_HOUR",
	}
	ymlConfig := YmlConfig{Server: serverConfig}

	// Test data
	os.Setenv("SERVER_PORT", "8080")
	os.Setenv("SERVER_PRIVATE_KEY", "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlKS2dJQkFBS0NBZ0VBd1BWbWFGdjJpZ0lvL1ZITk45dytoWUZDdHNFby8rZ0gwaWtpOEh3RlBybnBpVDN2Ci8wckptRTNyV2pBWExwU1V5NGh2R1dWUzVPMzlYcUgyd3dzcU5uYlNySUd6T3NKWGx4NVNOdy9BT0VGV012YzcKVysvQk9FY0xGUVFLdnJJNTlpUDZqZS91YitkY1F6NDc1bFNFNkpHbFczRlh1dkVJelQrYlliNEZYYTVrRDJJMAppZUVHc2hKV2RmNmFKQjMzeGZlYjI0WjUwSThmdXBsQlAvZnkzSnFJNDlEeGppTm5XSS9TWC9iMzdmRGppd3BqCkRVUjFBMzVxSTJ3VC94QlM3dWpEeG56c2o5N2JocFFWeXl5d2JET3AvNWNqZ3RuUGZqRHRXL2hndExNREVMOFMKWmpKekJXdXpPdlRzWWRNYkUrWU5mcDJLTmh3R0lCODVmc2tSWFFMUll6U2pwMHRrZ3dUNDFidTB2R2F4UC9QawpTUDJtcm9SM1ZMOXVidmJWazVjRUZScnNIVWR5TVYxWGpYYStidEo5bDV2WVp4NDF4OE1teWxGOExQc0V3cFpICk1VNXRXd0l4MWVXTEJvSVpWMjNRNklQMVV1TU50ZkYyT0dkQVlwWmJFM1VaMXp6K0lJeHFoMjhSN2xnVmcyaEwKbmtENE9zNy9GTGRuT2xDK2hDbGVyWXNTbXJibnNTUVBMVitrRHlYaUN4M3NRQmRHSVV6cmI4MnhWK05IaTBHQwpURFVQTGUyOGF6NS9pTXYyZHRJTzU1Mmk1SjN4WFU1YlNvYkVLdktyN09GQmdMUTUzY0FVZWk5S1Ezd3VPL1l6CkdrR1pLd1laR1pBMHRyTXpZcWNYdUYyMkk3ZFd1Ny9xUCtnSlJ1R0tvQitVSXJUUnZ4UllHYUdKRVNzQ0F3RUEKQVFLQ0FnQnE2QXVqQ2tDZjJlNkgyMGlPQ0hLRFdVaHpKTFhvZ3MvQ2VwUW5GUzk2djFwS2RZeUFyeGplVDExMApER3pybTlxTW9ieWNIMjA3OVRlSnRNYVk3Wmluc0ZHc3pmZFZPTk42b0l3QWdiT0g5M3NncGFXM25EQTdVL0VwCjVhRm1ZaXlHMlF3Nms1SlJZYWZXZ2FhQ1NKV1NuUTgyaUtlSXBYNnc2T3JYem9YK2oxNVV2NTVJUGVxTndtY3cKM0t6ZUVkUnpWR3AveUJPNW4zSisyZVl5NE5jbnJsN2xpYUFybGlYdUJWRVRaaXViSXZtTzBXUnJ0MS8xWFpiRQpwVjYyMUg1K2EydjhqcjRxUDlqMHlSc0NCSEtDb0dVeGZMVnFDKzlRQlIrYzh4SHhTN1VKRkRQSGQzQU9zc3NNCmdTSnVXSjRKK2szRHR0a2FmWkJQUDcyRkp6VjhrRHVEcDNSZnlzb1Q0eFBRUGtreFIwK1MvY2pYUXJjcnhzY2gKRHM1SFdwM2ZxTDh2dEtuRlJ2WlMrbm1HaVVYMmdLMXpnUmJLYWdvR3p6Zzl1V1gvaTl0Zk1hZmJyQWdMV21YcwpKblhBRFU1T2duRGdTN01TU1U2Rm5pdlZJSkdKNWJBNFNBRXA1OVdSRi92ckFlcWRDNXVlUkJtcVdtMzNYck1aCkRKUnNZb0hYU041dHlXdGFqaWNpMU1YclZnUkorWlh0cEhTVFBGUUdiSzEwWmZiYzdoQjZGMExXSGpsbngyaTYKcmtRUjVDZ2NxemRndkFCeU5kSmg2bnIyWitoVjFGeUNTbFgvWFZnMllUN1ZpQzdYNTJ4cWluTzNuN0dJOHFOQwpJL0VyZkNmS0pMUGJhTVdmQ2E0TEozaE9uT0d1SU9WdTZCV2VRRHNHU0JFSndjWFdzUUtDQVFFQS90QXp4R01tCjRsT2FSNGFaRkxQRzE0bjFyQUMrbU1OTU9pRzJJeXczSWdsRGVUZkI1S05oampBUWh2NHAyVnJraU9RZDQyOTQKWHBqdnBuSGw1MnZzelpwWlFKeW1zSGFLYjdJNHdCcndaUHBTSTUzd0hYUnpLNTFNQWZpdDhsQlJSWXJ0cWFhLwp6TlRISHdtcTJ0YlloY0krN0dab1NLcndReGUxWUQra1ZCK2REeXNkSkNhQ2VxaDdyQzYrUWdDRG1HeUxMejN2ClZDUWltS0gxaWpCNVJOTDJqM1BEVkwvUVQyYkZtb0FZWlI5ZmxuUUFDV1hPWERjRlpYSzY3SGdmQW13OXR6Z2kKaDU5REllaDZJWW5nWmJ1NzJnNmVyRHo3YVA3T3FrMVhzMHVrdDVVU3p6ZEZON1dXZWIraXdCdTNvZ3g2NFYxQQpCbWJ1bWtzNkN6ZXJsd0tDQVFFQXdkdHp6bkh4dnBML2t6NmgvbjAwVlJndm9vUU9oWTZGbVdDTWdXTWcwNWRECmlMdU9WbmRBazRObXVyVkRScXBYaFNYRFZqbFU4ZnBGSlJNNFRYRFpQaHJEUXlUNEp4L2xDeUxNYkd3RmtPb1cKYXNNZUdoVGdQKzNjc25BVVhCb0FDSmFFbENad3RBaEtYdXI4YnhQdlJMaHAxc2FHTjRjVDR4bVBGb2pOMlB5RQorRXFBWHRrNlhOMCtZVy9oZGFpZkdOL1BwRGFMcEF1MTJzVGFIR1NhOVBDVkNTbzdweHMveG9Tc0xTYW0rL3Y4Cm1lV1ZNVVZaT0RTbjA3a01rbXh3SWtNTXV5eVF5V0VzWWJMM3N1bFphYVBFMU1Ndnk3bUFKR0ppYklBekdERWkKYmJMSk9pcGc1TUMySHZTNFZYeFBGZGVTK3BWUTc4QllvRnBVajhySmpRS0NBUUVBM0RpVWhPWXNkTzVNS0FUcgp5RGlYWVRDYVlrMUNiRVJkWE9CRnlhQXRCZjE3a3dmZFN2enBFem4zRHJRYTl2N1hCSGdpWEsyNkdnZVRGd2JZCjYya2EvNWFtREhGV25xdlVlVFJPVjdqd2lsVE5LSHNYU2wyYUs5ZUdHUzRUSjVqQ3BKZXRUeklPRWJqVFhyKzgKS2VZRXU1VmxUR28xTnBpRmpYYXdDcjcyQnI1THZ4QkQzenBwQ2hrU3lYeWNjZTUvelB3Q1RwSDRoWCsxWnJTUwp3UnVqc3hlZ2Y5cE10cklRRm85N3VFdDh5ZWlUZERSTTA5Sm94c25Hb0NiSDVoYnF0ZTFXYVVMYWxOdlA2VDVDClR6b1o5ZEtLUjZyYTk0Qzh4OEZ3V3o3OHpMaFRZMVl6SzJOWkx3eUJRRGVmTU9qRGpBbTlLWWl1RE5wbzNIQ24KZVlwamdRS0NBUUVBdGxJMEgzU1EzU0NabUIxdTg4OURtY2lPZkhWZ3h3R2M2dnlnQ014M1FpbGdqY2QvL2hoWQpOcVI5eUpuVDlURWQ4UTdzSVRyNGhrQlFLYWRpNjRwMzl1M3F1VXFhelFrMVBIejA3Ly9FV0YrZ3g3Wk1xRkQ3Cis0UTFiZWoxYlEzUy9FQzczaTR0RDFWQXhQYVNoZEdrMWVmdk90MHB2QzJoYVpSUE8rMWNWSGhpZ3JabTkwMnMKazB4TmNBeHVhbDgxaW9wc1dsQW1reG1rWm1WL2tQYVp1a1pPbFBrUWM0Q3dRWC9rQXU3NFc4UEo5ZCt6cWt4RAp0aFhueGJ1amRFN2lRNGIyQVUvUHVHWlkvR1g2aWx6bkIvRExuU01aMzZ2T05lb0dFVytkSG1LUHM4WlRkUTRJClpQeE9ETjB5Uk13T0FVZm5aeDlwcUtNcGQxNmRhME5ZdlFLQ0FRRUFqU29QeWdTY1pCeUNLZ09GaDNiQ3U2czUKZGdsRUl3d3gvWWFuMEVkb0wySkpzVUl1c3JrelFHZGpXODFoSmJFY09vMmRqc24wK3FUbUlBMEdnZmRXTTZFUAo0VUFYUzRSZFZObWR5K0JyNXdiUlQ3YmE5V1JVS0JFY0xGZzFsUEx2SXVsTkNJZ3k0cHlKTy9ENEVqeWNIK1E3CkQrOUVIOVRpTXFXTytCOENrU3pIMFUxN04vR0Z2RHFXU25FWDNXWVZDMS92aG1DWXJQZENYMTVCY3BLV0NkdjAKbW1IaEtJZGl3OWdIaUlPS2daWHVRanZRUW1OT0p5R0ZlMysvT2UycUJlbTVVQ0lhZGhTVFNTcDZ2dzdNZm9teApOWnNCVkg1cS9oV3Q1ZlgyV3ZoZFBIZlJFYTVFYnlUbThPZGNJc0RoN1VpVUJ0dU8rYkFDM2VBRWZIcUZlUT09Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==")
	os.Setenv("SERVER_TOKEN_EXPIRE_HOUR", "100")

	if !strings.EqualFold(ymlConfig.GetServerConfig().Port, "8080") {
		t.Errorf("Incorrect GetServerConfig test. port = %s", ymlConfig.GetServerConfig().Port)
		t.FailNow()
	}

	if !strings.EqualFold(ymlConfig.GetServerConfig().SignedInPrivateKeyBase64, "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlKS2dJQkFBS0NBZ0VBd1BWbWFGdjJpZ0lvL1ZITk45dytoWUZDdHNFby8rZ0gwaWtpOEh3RlBybnBpVDN2Ci8wckptRTNyV2pBWExwU1V5NGh2R1dWUzVPMzlYcUgyd3dzcU5uYlNySUd6T3NKWGx4NVNOdy9BT0VGV012YzcKVysvQk9FY0xGUVFLdnJJNTlpUDZqZS91YitkY1F6NDc1bFNFNkpHbFczRlh1dkVJelQrYlliNEZYYTVrRDJJMAppZUVHc2hKV2RmNmFKQjMzeGZlYjI0WjUwSThmdXBsQlAvZnkzSnFJNDlEeGppTm5XSS9TWC9iMzdmRGppd3BqCkRVUjFBMzVxSTJ3VC94QlM3dWpEeG56c2o5N2JocFFWeXl5d2JET3AvNWNqZ3RuUGZqRHRXL2hndExNREVMOFMKWmpKekJXdXpPdlRzWWRNYkUrWU5mcDJLTmh3R0lCODVmc2tSWFFMUll6U2pwMHRrZ3dUNDFidTB2R2F4UC9QawpTUDJtcm9SM1ZMOXVidmJWazVjRUZScnNIVWR5TVYxWGpYYStidEo5bDV2WVp4NDF4OE1teWxGOExQc0V3cFpICk1VNXRXd0l4MWVXTEJvSVpWMjNRNklQMVV1TU50ZkYyT0dkQVlwWmJFM1VaMXp6K0lJeHFoMjhSN2xnVmcyaEwKbmtENE9zNy9GTGRuT2xDK2hDbGVyWXNTbXJibnNTUVBMVitrRHlYaUN4M3NRQmRHSVV6cmI4MnhWK05IaTBHQwpURFVQTGUyOGF6NS9pTXYyZHRJTzU1Mmk1SjN4WFU1YlNvYkVLdktyN09GQmdMUTUzY0FVZWk5S1Ezd3VPL1l6CkdrR1pLd1laR1pBMHRyTXpZcWNYdUYyMkk3ZFd1Ny9xUCtnSlJ1R0tvQitVSXJUUnZ4UllHYUdKRVNzQ0F3RUEKQVFLQ0FnQnE2QXVqQ2tDZjJlNkgyMGlPQ0hLRFdVaHpKTFhvZ3MvQ2VwUW5GUzk2djFwS2RZeUFyeGplVDExMApER3pybTlxTW9ieWNIMjA3OVRlSnRNYVk3Wmluc0ZHc3pmZFZPTk42b0l3QWdiT0g5M3NncGFXM25EQTdVL0VwCjVhRm1ZaXlHMlF3Nms1SlJZYWZXZ2FhQ1NKV1NuUTgyaUtlSXBYNnc2T3JYem9YK2oxNVV2NTVJUGVxTndtY3cKM0t6ZUVkUnpWR3AveUJPNW4zSisyZVl5NE5jbnJsN2xpYUFybGlYdUJWRVRaaXViSXZtTzBXUnJ0MS8xWFpiRQpwVjYyMUg1K2EydjhqcjRxUDlqMHlSc0NCSEtDb0dVeGZMVnFDKzlRQlIrYzh4SHhTN1VKRkRQSGQzQU9zc3NNCmdTSnVXSjRKK2szRHR0a2FmWkJQUDcyRkp6VjhrRHVEcDNSZnlzb1Q0eFBRUGtreFIwK1MvY2pYUXJjcnhzY2gKRHM1SFdwM2ZxTDh2dEtuRlJ2WlMrbm1HaVVYMmdLMXpnUmJLYWdvR3p6Zzl1V1gvaTl0Zk1hZmJyQWdMV21YcwpKblhBRFU1T2duRGdTN01TU1U2Rm5pdlZJSkdKNWJBNFNBRXA1OVdSRi92ckFlcWRDNXVlUkJtcVdtMzNYck1aCkRKUnNZb0hYU041dHlXdGFqaWNpMU1YclZnUkorWlh0cEhTVFBGUUdiSzEwWmZiYzdoQjZGMExXSGpsbngyaTYKcmtRUjVDZ2NxemRndkFCeU5kSmg2bnIyWitoVjFGeUNTbFgvWFZnMllUN1ZpQzdYNTJ4cWluTzNuN0dJOHFOQwpJL0VyZkNmS0pMUGJhTVdmQ2E0TEozaE9uT0d1SU9WdTZCV2VRRHNHU0JFSndjWFdzUUtDQVFFQS90QXp4R01tCjRsT2FSNGFaRkxQRzE0bjFyQUMrbU1OTU9pRzJJeXczSWdsRGVUZkI1S05oampBUWh2NHAyVnJraU9RZDQyOTQKWHBqdnBuSGw1MnZzelpwWlFKeW1zSGFLYjdJNHdCcndaUHBTSTUzd0hYUnpLNTFNQWZpdDhsQlJSWXJ0cWFhLwp6TlRISHdtcTJ0YlloY0krN0dab1NLcndReGUxWUQra1ZCK2REeXNkSkNhQ2VxaDdyQzYrUWdDRG1HeUxMejN2ClZDUWltS0gxaWpCNVJOTDJqM1BEVkwvUVQyYkZtb0FZWlI5ZmxuUUFDV1hPWERjRlpYSzY3SGdmQW13OXR6Z2kKaDU5REllaDZJWW5nWmJ1NzJnNmVyRHo3YVA3T3FrMVhzMHVrdDVVU3p6ZEZON1dXZWIraXdCdTNvZ3g2NFYxQQpCbWJ1bWtzNkN6ZXJsd0tDQVFFQXdkdHp6bkh4dnBML2t6NmgvbjAwVlJndm9vUU9oWTZGbVdDTWdXTWcwNWRECmlMdU9WbmRBazRObXVyVkRScXBYaFNYRFZqbFU4ZnBGSlJNNFRYRFpQaHJEUXlUNEp4L2xDeUxNYkd3RmtPb1cKYXNNZUdoVGdQKzNjc25BVVhCb0FDSmFFbENad3RBaEtYdXI4YnhQdlJMaHAxc2FHTjRjVDR4bVBGb2pOMlB5RQorRXFBWHRrNlhOMCtZVy9oZGFpZkdOL1BwRGFMcEF1MTJzVGFIR1NhOVBDVkNTbzdweHMveG9Tc0xTYW0rL3Y4Cm1lV1ZNVVZaT0RTbjA3a01rbXh3SWtNTXV5eVF5V0VzWWJMM3N1bFphYVBFMU1Ndnk3bUFKR0ppYklBekdERWkKYmJMSk9pcGc1TUMySHZTNFZYeFBGZGVTK3BWUTc4QllvRnBVajhySmpRS0NBUUVBM0RpVWhPWXNkTzVNS0FUcgp5RGlYWVRDYVlrMUNiRVJkWE9CRnlhQXRCZjE3a3dmZFN2enBFem4zRHJRYTl2N1hCSGdpWEsyNkdnZVRGd2JZCjYya2EvNWFtREhGV25xdlVlVFJPVjdqd2lsVE5LSHNYU2wyYUs5ZUdHUzRUSjVqQ3BKZXRUeklPRWJqVFhyKzgKS2VZRXU1VmxUR28xTnBpRmpYYXdDcjcyQnI1THZ4QkQzenBwQ2hrU3lYeWNjZTUvelB3Q1RwSDRoWCsxWnJTUwp3UnVqc3hlZ2Y5cE10cklRRm85N3VFdDh5ZWlUZERSTTA5Sm94c25Hb0NiSDVoYnF0ZTFXYVVMYWxOdlA2VDVDClR6b1o5ZEtLUjZyYTk0Qzh4OEZ3V3o3OHpMaFRZMVl6SzJOWkx3eUJRRGVmTU9qRGpBbTlLWWl1RE5wbzNIQ24KZVlwamdRS0NBUUVBdGxJMEgzU1EzU0NabUIxdTg4OURtY2lPZkhWZ3h3R2M2dnlnQ014M1FpbGdqY2QvL2hoWQpOcVI5eUpuVDlURWQ4UTdzSVRyNGhrQlFLYWRpNjRwMzl1M3F1VXFhelFrMVBIejA3Ly9FV0YrZ3g3Wk1xRkQ3Cis0UTFiZWoxYlEzUy9FQzczaTR0RDFWQXhQYVNoZEdrMWVmdk90MHB2QzJoYVpSUE8rMWNWSGhpZ3JabTkwMnMKazB4TmNBeHVhbDgxaW9wc1dsQW1reG1rWm1WL2tQYVp1a1pPbFBrUWM0Q3dRWC9rQXU3NFc4UEo5ZCt6cWt4RAp0aFhueGJ1amRFN2lRNGIyQVUvUHVHWlkvR1g2aWx6bkIvRExuU01aMzZ2T05lb0dFVytkSG1LUHM4WlRkUTRJClpQeE9ETjB5Uk13T0FVZm5aeDlwcUtNcGQxNmRhME5ZdlFLQ0FRRUFqU29QeWdTY1pCeUNLZ09GaDNiQ3U2czUKZGdsRUl3d3gvWWFuMEVkb0wySkpzVUl1c3JrelFHZGpXODFoSmJFY09vMmRqc24wK3FUbUlBMEdnZmRXTTZFUAo0VUFYUzRSZFZObWR5K0JyNXdiUlQ3YmE5V1JVS0JFY0xGZzFsUEx2SXVsTkNJZ3k0cHlKTy9ENEVqeWNIK1E3CkQrOUVIOVRpTXFXTytCOENrU3pIMFUxN04vR0Z2RHFXU25FWDNXWVZDMS92aG1DWXJQZENYMTVCY3BLV0NkdjAKbW1IaEtJZGl3OWdIaUlPS2daWHVRanZRUW1OT0p5R0ZlMysvT2UycUJlbTVVQ0lhZGhTVFNTcDZ2dzdNZm9teApOWnNCVkg1cS9oV3Q1ZlgyV3ZoZFBIZlJFYTVFYnlUbThPZGNJc0RoN1VpVUJ0dU8rYkFDM2VBRWZIcUZlUT09Ci0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==") {
		t.Errorf("Incorrect GetServerConfig test. privaet_key_base64 = %s", ymlConfig.GetServerConfig().SignedInPrivateKeyBase64)
		t.FailNow()
	}

	if !strings.EqualFold(ymlConfig.GetServerConfig().TokenExpireHourStr, "100") {
		t.Errorf("Incorrect GetServerConfig test. token-expire-hour = %s", ymlConfig.GetServerConfig().TokenExpireHourStr)
		t.FailNow()
	}
}

// GetEtcdConfig test
func TestGetEtcdConfig(t *testing.T) {
	etcdConfig := EtcdConfig{Host: "$ETCD_HOST", Port: "$ETCD_PORT"}
	ymlConfig := YmlConfig{Etcd: etcdConfig}

	// Test data
	os.Setenv("ETCD_HOST", "localhost")
	os.Setenv("ETCD_PORT", "2380")

	if !strings.EqualFold(ymlConfig.GetEtcdConfig().Host, "localhost") {
		t.Errorf("Incorrect GetEtcdConfig test. host = %s", ymlConfig.GetEtcdConfig().Host)
		t.FailNow()
	}

	if !strings.EqualFold(ymlConfig.GetEtcdConfig().Port, "2380") {
		t.Errorf("Incorrect GetEtcdConfig test. port = %s", ymlConfig.GetEtcdConfig().Port)
		t.FailNow()
	}
}

// GetDbConfig test
func TestGetDbConfig(t *testing.T) {
	dbConfig := DbConfig{
		Engine:   "$DB_ENGINE",
		Host:     "$DB_HOST",
		User:     "$DB_USER",
		Password: "$DB_PASSWORD",
		Port:     "$DB_PORT",
		Name:     "$DB_NAME",
	}
	ymlConfig := YmlConfig{Db: dbConfig}

	// Test data
	os.Setenv("DB_ENGINE", "mysql")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASSWORD", "root")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_NAME", "grant_n_z")

	if !strings.EqualFold(ymlConfig.GetDbConfig().Engine, "mysql") {
		t.Errorf("Incorrect GetEtcdConfig test. engine = %s", ymlConfig.GetDbConfig().Engine)
		t.FailNow()
	}

	if !strings.EqualFold(ymlConfig.GetDbConfig().Host, "localhost") {
		t.Errorf("Incorrect GetEtcdConfig test. host = %s", ymlConfig.GetDbConfig().Host)
		t.FailNow()
	}

	if !strings.EqualFold(ymlConfig.GetDbConfig().User, "root") {
		t.Errorf("Incorrect GetEtcdConfig test. user = %s", ymlConfig.GetDbConfig().User)
		t.FailNow()
	}

	if !strings.EqualFold(ymlConfig.GetDbConfig().Password, "root") {
		t.Errorf("Incorrect GetEtcdConfig test. password = %s", ymlConfig.GetDbConfig().Password)
		t.FailNow()
	}

	if !strings.EqualFold(ymlConfig.GetDbConfig().Port, "3306") {
		t.Errorf("Incorrect GetEtcdConfig test. port = %s", ymlConfig.GetDbConfig().Port)
		t.FailNow()
	}

	if !strings.EqualFold(ymlConfig.GetDbConfig().Name, "grant_n_z") {
		t.Errorf("Incorrect GetEtcdConfig test. db = %s", ymlConfig.GetDbConfig().Name)
		t.FailNow()
	}
}
