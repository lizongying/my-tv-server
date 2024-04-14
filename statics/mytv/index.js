import {decode, encode, verify} from './gua64.min.js';

const input = document.getElementById('input');
const local = document.getElementById('local');
const lan = document.getElementById('lan');
const internet = document.getElementById('internet');
const localJson = document.getElementById('local-json');
const lanJson = document.getElementById('lan-json');
const internetJson = document.getElementById('internet-json');
const options = document.querySelectorAll('.option');
const messages = document.querySelectorAll('.message');

const allowedExtensions = ['json', 'txt', 'm3u'];

const nameRegex = /tvg-name="([^"]+)"/;
const logRegex = /tvg-logo="([^"]+)"/;
const groupRegex = /group-title="([^"]+)"/;

let outType = 'raw';

let files = null;

options.forEach(option => {
    option.addEventListener('click', function () {
        if (option.dataset.type !== outType) {
            local.href = ''
            local.innerText = ''
            lan.href = ''
            lan.innerText = ''
            internet.href = ''
            internet.innerText = ''
            if (files != null) {
                handleFiles(files)
            }
        }
        outType = option.dataset.type

        options.forEach(opt => opt.classList.remove('active'));

        this.classList.add('active');

        let id = this.dataset.message;
        messages.forEach(message => {
            if (message.id === id) {
                message.classList.remove('hide')
            } else {
                message.classList.add('hide')
            }
        });
    });
});

input.onchange = function (e) {
    files = e.target.files;
    handleFiles(files)
}

const handleFiles = async function (files) {
    for (let i = 0; i < files.length; i++) {
        const file = files[i];
        const extension = file.name.split('.').pop().toLowerCase();
        if (!allowedExtensions.includes(extension)) {
            return;
        }
        await read(file, extension);
    }
}

const read = async (file, extension) => {
    const reader = new FileReader();
    reader.onload = (e) => {
        let out = e.target.result
        switch (outType) {
            case 'raw': {
                break
            }
            case 'json': {
                if (verify(out)) {
                    out = decode(out)
                }
                switch (extension) {
                    case 'txt': {
                        const lines = out.split('\n');
                        const list = [];
                        let group = '';
                        for (let line of lines) {
                            line = line.trim();
                            if (line !== '') {
                                if (line.includes('#genre#')) {
                                    group = line.split(',', 1)[0].trim();
                                } else {
                                    const arr = line.split(',').map(v => v.trim())
                                    list.push({
                                        group: group,
                                        title: arr.shift(),
                                        uris: arr
                                    })
                                }
                            }
                        }
                        out = JSON.stringify(list)
                        break
                    }
                    case 'm3u': {
                        const lines = out.split('\n');
                        const list = [];
                        for (let line of lines) {
                            line = line.trim();
                            if (line.startsWith('#EXTINF')) {
                                const infos = line.split(',');
                                const title = infos.pop()
                                list.push({
                                    name: infos[0]?.match(nameRegex)[1],
                                    group: infos[0]?.match(groupRegex)[1],
                                    logo: infos[0]?.match(logRegex)[1],
                                    title: title,
                                    uris: [lines[lines.indexOf(line) + 1]]
                                })
                            }
                        }
                        out = JSON.stringify(list)
                        break
                    }
                }
                let obj = JSON.parse(out);
                console.log(obj);
                break
            }
            case 'txt': {
                break
            }
            case 'm3u': {
                break
            }
        }

        sendFile(out);
    };
    reader.readAsText(file);
}

const sendFile = async (body) => {
    const response = await fetch('/upload', {
        method: 'POST',
        body: body
    });
    const json = await response.json();
    console.log(json);
    const localUrl = json?.data?.local_url
    if (localUrl) {
        local.href = `http://${localUrl}`;
        local.innerText = localUrl;
    }
    const lanUrl = json?.data?.lan_url
    if (lanUrl) {
        lan.href = `http://${lanUrl}`;
        lan.innerText = lanUrl;
    }
    const internetUrl = json?.data?.internet_url
    if (internetUrl) {
        internet.href = `http://${internetUrl}`;
        internet.innerText = internetUrl;
    }
}

(async ()=>{
    const response = await fetch('/ip/info');
    const json = await response.json();
    const localUrl = json?.data?.local_url
    if (localUrl) {
        localJson.href = `http://${localUrl}/example.json`;
        localJson.innerText = `http://${localUrl}/example.json`;
    }
    const lanUrl = json?.data?.lan_url
    if (lanUrl) {
        lanJson.href = `http://${lanUrl}/example.json`;
        lanJson.innerText = `http://${lanUrl}/example.json`;
    }
    const internetUrl = json?.data?.internet_url
    if (internetUrl) {
        internetJson.href = `http://${internetUrl}/example.json`;
        internetJson.innerText = `http://${internetUrl}/example.json`;
    }
})()