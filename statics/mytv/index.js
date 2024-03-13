const input = document.getElementById('input');
const local = document.getElementById('local');
const lan = document.getElementById('lan');
const internet = document.getElementById('internet');

input.onchange = async function handleFiles(e) {
    const files = e.target.files;
    for (let i = 0; i < files.length; i++) {
        const file = files[i];
        if (!file.name.endsWith('.json')) {
            return;
        }
        await read(file);
    }
}

const read = async (file) => {
    const reader = new FileReader();
    reader.onload = (e) => {
        let obj = JSON.parse(e.target.result);
        console.log(obj);
        sendFile(obj);
    };
    reader.readAsText(file);
}

const sendFile = async (body) => {
    const response = await fetch('/upload', {
        method: 'POST',
        body: JSON.stringify(body)
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