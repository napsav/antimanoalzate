/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Wait for the deviceready event before using any of Cordova's device APIs.
// See https://cordova.apache.org/docs/en/latest/cordova/events/events.html#deviceready
document.addEventListener('deviceready', onDeviceReady, false);

function onDeviceReady() {
    // Cordova is now initialized. Have fun!

    console.log('Running cordova-' + cordova.platformId + '@' + cordova.version);
    document.querySelector('body').classList.remove('hidden');
}

document.getElementById("save").addEventListener('click', salvaImpostazioni, false)
let pronto = false;
let ipSalvato = "";
let portaSalvata = "8080";
function salvaImpostazioni() {
		let errors = [];
		let ip = document.getElementById('indirizzoip').value
		let porta = document.getElementById('porta').value
		let ipPattern = new RegExp(/^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/)
		if(ipPattern.test(ip)) {
			document.getElementById('pronto').classList.remove('hidden')
			pronto = true
			ipSalvato = ip
		} else {
			errors.push("Indirizzo IP non valido")
		}
		if (errors.length > 0) {
			for(elem of errors) {
				let err = document.createElement("p")
				err.innerHTML = elem
				err.style="color:red"
				document.body.append(err)
			}
		}
}


let log = document.getElementById("log")

	document.addEventListener("volumeupbutton", volumeUp, false)
	document.addEventListener("volumedownbutton", volumeDown, false)


function httpGetAsync(theUrl)
{
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.onreadystatechange = function() { 
        if (xmlHttp.readyState == 4 && xmlHttp.status == 200)
            console.log(xmlHttp.responseText);
    }
    xmlHttp.open("GET", theUrl, true); // true for asynchronous 
    xmlHttp.send(null);
}

function volumeUp() {
	log.innerHTML += "volume sopra\n"
	httpGetAsync("http://"+ipSalvato+":"+portaSalvata+"/up")
}
function volumeDown() {
	log.innerHTML += "volume sotto\n"
	httpGetAsync("http://"+ipSalvato+":"+portaSalvata+"/down")
}

function schermoAttivo() {
	window.plugins.insomnia.keepAwake()
	document.getElementById('schermoAttivo').classList.add('hidden')
	document.getElementById('schermoNormale').classList.remove('hidden')
}


function schermoNormale() {
	window.plugins.insomnia.allowSleepAgain()
	document.getElementById('schermoAttivo').classList.remove('hidden')
	document.getElementById('schermoNormale').classList.add('hidden')
}

function nascondi() {
    document.getElementById('touch').classList.remove('hidden')
    document.getElementById('normale').classList.add('hidden')
}