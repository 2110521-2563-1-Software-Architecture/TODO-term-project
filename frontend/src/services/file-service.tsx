import axios from 'axios'

export async function uploadFile(formdata: FormData, onUploadProgress: any) {
    return axios.post('localhost:9000/uploadFile/', formdata, {
        onUploadProgress
    })
}

export async function getAvailableFiles() {
    return axios.get('url/files') // to be replaced with actual url
}
