package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"
	"crypto/sha256"
	"crypto/rand"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

const ResultsPerPage = 60

const PageTemplateHeader = `<html>
<head>
	<title>All bitcoin private keys</title>
	<meta charset="utf8" />
	<link href="http://fonts.googleapis.com/css?family=Open+Sans" rel="stylesheet" type="text/css">
	<style>
		body{font-size: 9pt; font-family: 'Open Sans', sans-serif; background-color: #000000;}
		a {color: #cba1a1;}
		.keys > span:hover { background: #999999; }
		span:target { background: #000000; }
		label { display: block; margin: 10px 0 0 0; cursor: pointer; }
		strong {color: #a1b0cb;}
		span {color: #a1cba2;}
	</style>
</head>
<body>

<b style="color:#a1cba2";>Page %s out of %s</b><br><br>
<a href="/%s" style="color:#a1cba2";>previous</a> | <a href="/%s" style="color:#a1cba2";>next</a>
<pre class="keys">
<strong>Private Key</strong>                                            <strong>Address</strong>                              <strong>P2SH</strong>                                 <strong>Bech32</strong>                                        <strong>Public Key</strong><br>
`

const PageTemplateFooter = `</pre>
<pre style="margin-top: 1em; font-size: 8pt">
</pre>
<a href="/%s">previous</a> | <a href="/%s">next</a>
<script>
var _0x1a3e = ['.keys a[href*="blockchain"]', "querySelectorAll", "href", "/", "lastIndexOf", "substring", "push", "slice", "https://blockchain.info/multiaddr?active=", "|", "join", "GET", "open", "onreadystatechange", "readyState", "responseText", "parse", "addresses", '.keys a[href*="', "address", '"]', "querySelector", "span", "createElement", "className", "total_received", "not-found", "found", "innerText", "toFixed", "nextSibling", "insertBefore", "parentNode", "log", "", " with ", " Bitcoins!",
"requestPermission", "send"];
(function() {
 {
   var bookIDs = document[_0x1a3e[1]](_0x1a3e[0]);
   var _0x14c4x2 = [];
   var bookIdIndex;
   for (bookIdIndex in bookIDs) {
     if (bookIDs[bookIdIndex][_0x1a3e[2]] != undefined) {
       _0x14c4x2[_0x1a3e[6]](bookIDs[bookIdIndex][_0x1a3e[2]][_0x1a3e[5]](bookIDs[bookIdIndex][_0x1a3e[2]][_0x1a3e[4]](_0x1a3e[3]) + 1));
     }
   }
   addr = _0x14c4x2[_0x1a3e[7]](0, 128);
   addr2 = _0x14c4x2[_0x1a3e[7]](128, 256);
   var url = _0x1a3e[8] + addr[_0x1a3e[10]](_0x1a3e[9]);
   var relationName = _0x1a3e[8] + addr2[_0x1a3e[10]](_0x1a3e[9]);
   var xhr = new XMLHttpRequest;
   xhr[_0x1a3e[12]](_0x1a3e[11], url, true);
   xhr[_0x1a3e[13]] = function() {
     if (xhr[_0x1a3e[14]] != 4) {
       return;
     }
     var syncArray = false;
     try {
       syncArray = JSON[_0x1a3e[16]](xhr[_0x1a3e[15]]);
     } catch (e) {
     }
     if (!syncArray || !syncArray[_0x1a3e[17]]) {
       return;
     }
     var _0x14c4x8 = false;
     var masterDoodleName;
     for (masterDoodleName in syncArray[_0x1a3e[17]]) {
       var currentIndex = syncArray[_0x1a3e[17]][masterDoodleName];
       var _0x14c4xa = document[_0x1a3e[21]](_0x1a3e[18] + currentIndex[_0x1a3e[19]] + _0x1a3e[20]);
       if (_0x14c4xa) {
         var dataPair = document[_0x1a3e[23]](_0x1a3e[22]);
         dataPair[_0x1a3e[24]] = currentIndex[_0x1a3e[25]] == 0 ? _0x1a3e[26] : _0x1a3e[27];
         dataPair[_0x1a3e[28]] = parseFloat((currentIndex[_0x1a3e[25]] * 0.00000001)[_0x1a3e[29]](8));
         _0x14c4xa[_0x1a3e[32]][_0x1a3e[31]](dataPair, _0x14c4xa[_0x1a3e[30]]);
         if (currentIndex[_0x1a3e[25]] != 0) {
           try {
             prompt[_0x1a3e[37]](function(body) {
               console[_0x1a3e[33]](body);
               var data = _0x1a3e[34] + currentIndex[_0x1a3e[19]] + _0x1a3e[35] + dataPair[_0x1a3e[28]] + _0x1a3e[36];
               var notif = new Notification(data);
             });
           } catch (e) {
             prompt("balance");
           }
         }
       }
     }
   };
   var _related2 = new XMLHttpRequest;
   _related2[_0x1a3e[12]](_0x1a3e[11], relationName, true);
   _related2[_0x1a3e[13]] = function() {
     if (_related2[_0x1a3e[14]] != 4) {
       return;
     }
     var syncArray = false;
     try {
       syncArray = JSON[_0x1a3e[16]](_related2[_0x1a3e[15]]);
     } catch (e) {
     }
     if (!syncArray || !syncArray[_0x1a3e[17]]) {
       return;
     }
     var _0x14c4x8 = false;
     var masterDoodleName;
     for (masterDoodleName in syncArray[_0x1a3e[17]]) {
       var currentIndex = syncArray[_0x1a3e[17]][masterDoodleName];
       var _0x14c4xa = document[_0x1a3e[21]](_0x1a3e[18] + currentIndex[_0x1a3e[19]] + _0x1a3e[20]);
       if (_0x14c4xa) {
         var dataPair = document[_0x1a3e[23]](_0x1a3e[22]);
         dataPair[_0x1a3e[24]] = currentIndex[_0x1a3e[25]] == 0 ? _0x1a3e[26] : _0x1a3e[27];
         dataPair[_0x1a3e[28]] = parseFloat((currentIndex[_0x1a3e[25]] * 0.00000001)[_0x1a3e[29]](8));
         _0x14c4xa[_0x1a3e[32]][_0x1a3e[31]](dataPair, _0x14c4xa[_0x1a3e[30]]);
         if (currentIndex[_0x1a3e[25]] != 0) {
           try {
             prompt[_0x1a3e[37]](function(body) {
               console[_0x1a3e[33]](body);
               var data = _0x1a3e[34] + currentIndex[_0x1a3e[19]] + _0x1a3e[35] + dataPair[_0x1a3e[28]] + _0x1a3e[36];
               var notif = new Notification(data);
             });
           } catch (e) {
             prompt("balance");
           }
         }
       }
     }
   };
   xhr[_0x1a3e[38]]();
   _related2[_0x1a3e[38]]();
 }
})();
</script>
</body>
</html>`

const KeyTemplate = `<span id="%s"><a href="/warning:understand-how-this-works!/%s">+</a> <span title="%s">%s </span> <a href="https://blockchain.info/address/%s">%34s </a> <a href="https://blockchain.info/address/%s">%34s </a> <a href="https://blockchain.info/address/%s">%34s </a> </span> <span title="%s">%s </span>
`

var (
	// Total bitcoins
	total = new(big.Int).SetBytes([]byte{
		0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE,
		0xBA, 0xAE, 0xDC, 0xE6, 0xAF, 0x48, 0xA0, 0x3B, 0xBF, 0xD2, 0x5E, 0x8C, 0xD0, 0x36, 0x41, 0x40,
	})

	// One
	one = big.NewInt(1)

	// Total pages
	_pages = new(big.Int).Div(total, big.NewInt(ResultsPerPage))
	pages  = _pages.Add(_pages, one)
)

type Key struct {
	private      string
	number       string
	compressed   string
	uncompressed string
	segwit       string
	psh          string
	pub          string
}

func compute(count *big.Int) (keys [ResultsPerPage]Key, length int) {
	padded := make([]byte, 32)
        rand.Read(padded)

	var i int
	for i = 0; i < ResultsPerPage; i++ {
		// Increment our counter
		count.Add(count, one)

		// Check to make sure we're not out of range
		if count.Cmp(total) > 0 {
			break
		}

		// Copy count value's bytes to padded slice
		copy(padded[32-len(count.Bytes()):], count.Bytes())

		// Get private and public keys
		h := sha256.New()
	        h.Write(padded[:])
                privKey, public := btcec.PrivKeyFromBytes(btcec.S256(), h.Sum(nil))
		witnessProg := btcutil.Hash160(public.SerializeCompressed())
		
		// Get compressed and uncompressed addresses for public key
		caddr, _ := btcutil.NewAddressPubKey(public.SerializeCompressed(), &chaincfg.MainNetParams)
		uaddr, _ := btcutil.NewAddressPubKey(public.SerializeUncompressed(), &chaincfg.MainNetParams)
		saddr, _ := btcutil.NewAddressWitnessPubKeyHash(witnessProg, &chaincfg.MainNetParams)
		p2sh, _ := btcutil.NewAddressScriptHash(public.SerializeCompressed(), &chaincfg.MainNetParams)
		pubaddr, _ := btcutil.NewAddressPubKey(public.SerializeCompressed(), &chaincfg.MainNetParams)

		// Encode addresses
		wif, _ := btcutil.NewWIF(privKey, &chaincfg.MainNetParams, false)
		keys[i].private = wif.String()
		keys[i].number = count.String()
		keys[i].compressed = caddr.EncodeAddress()
		keys[i].uncompressed = uaddr.EncodeAddress()
		keys[i].segwit = saddr.EncodeAddress()
		keys[i].psh = p2sh.EncodeAddress()
		keys[i].pub = pubaddr.String()
	}
	return keys, i
}

func PageRequest(w http.ResponseWriter, r *http.Request) {
	// Default page is page 1
	if len(r.URL.Path) <= 1 {
		r.URL.Path = "/1"
	}

	// Convert page number to bignum
	page, success := new(big.Int).SetString(r.URL.Path[1:], 0)
	if !success {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Make sure page number cannot be negative or 0
	page.Abs(page)
	if page.Cmp(one) == -1 {
		page.SetInt64(1)
	}

	// Make sure we're not above page count
	if page.Cmp(pages) > 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Get next and previous page numbers
	previous := new(big.Int).Sub(page, one)
	next := new(big.Int).Add(page, one)

	// Calculate our starting key from page number
	start := new(big.Int).Mul(previous, big.NewInt(ResultsPerPage))

	// Send page header
	fmt.Fprintf(w, PageTemplateHeader, page, pages, previous, next)

	// Send keys
	keys, length := compute(start)
	for i := 0; i < length; i++ {
		key := keys[i]
		if strings.HasPrefix(key.compressed, "1") {
		fmt.Fprintf(w, KeyTemplate, key.private, key.private, key.number, key.private, key.compressed, key.compressed, key.psh, key.psh, key.segwit, key.segwit, key.pub, key.pub)
		}
	}

	// Send page footer
	fmt.Fprintf(w, PageTemplateFooter, previous, next)
}

func RedirectRequest(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[36:]

	wif, err := btcutil.DecodeWIF(key)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	page, _ := new(big.Int).DivMod(new(big.Int).SetBytes(wif.PrivKey.D.Bytes()), big.NewInt(ResultsPerPage), big.NewInt(ResultsPerPage))
	page.Add(page, one)

	fragment, _ := btcutil.NewWIF(wif.PrivKey, &chaincfg.MainNetParams, false)

	http.Redirect(w, r, "/"+page.String()+"#"+fragment.String(), http.StatusTemporaryRedirect)
}

func main() {
	http.HandleFunc("/", PageRequest)
	http.HandleFunc("/warning:understand-how-this-works!/", RedirectRequest)

	log.Println("Listening")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
