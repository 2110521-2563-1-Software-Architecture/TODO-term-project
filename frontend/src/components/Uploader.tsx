import React from 'react'
import { uploadFile } from '../services/file-service'

interface IProps {

}

interface UploaderState {
    selectedFile?: any;
    progress?: number;
}

export default class Uploader extends React.Component<IProps, UploaderState> {
    constructor(props: any) {
        super(props)

        this.state = {
            selectedFile: undefined,
            progress: 0,
        }
    }

    onSelectFile = (event: any) => {
        this.setState({
            selectedFile: event.target.files[0]
        })
    }

    onFileUpload = async () => {
        this.setState({progress: 0})

        const formdoc: any = document.getElementById('upload-form')
        const formdata: FormData = new FormData(formdoc)

        try {
            const uploadResponse = await uploadFile(formdata, (event: any) => {
                this.setState({
                    progress: (Math.round((100 * event.loaded) / event.total))
                })
            })

            alert('File upload completed.')
        }
        catch (e) {
            this.setState({
                progress: 0,
            })
            alert('Could not upload the file!')
        }

        this.setState({
            selectedFile: undefined
        })
    }

    render() {
        return (
            <div>
                <h2>
                    Upload file here
                </h2>
                <div>
                    <form id="upload-form" name="upload-form">
                        <div>
                            <input type="file" id="file" name="file" onChange={this.onSelectFile} />
                            <button onClick={this.onFileUpload}>
                                Upload
                            </button>
                        </div>
                    </form>
                </div>
                <br />
                <div>
                    Upload progress : {this.state.progress}%
                </div>
            </div>
        )
    }
}
