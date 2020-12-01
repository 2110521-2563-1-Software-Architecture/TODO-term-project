import axios from 'axios'

export async function uploadFile(formdata: FormData, onUploadProgress: any) {
    return axios.post('url/upload', formdata, {
        onUploadProgress
    })
}

export async function getAvailableFiles() {
    return axios.get('url/files')
}

export async function downloadFile(filename: string) {
    return axios.get(`url/file/${filename}`)
}
