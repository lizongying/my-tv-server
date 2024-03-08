const input = document.getElementById('input');
const lan = document.getElementById('lan');

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
    const url = json?.data?.channels_url
    if (url) {
        lan.href = `http://${url}`;
        lan.innerText = url;
    }
}