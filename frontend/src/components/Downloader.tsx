import React from 'react'
import { getAvailableFiles } from '../services/file-service'

interface IProps {

}

interface DownloaderState {
    files?: any
}

export default class Downloader extends React.Component<IProps,DownloaderState> {
    constructor(props: any) {
        super(props)
    }

    async componentDidMount() {
        const files = await getAvailableFiles();
        this.setState({
            files
        })
    }

    generateFileData(fileinfo: any) {
        return (
            <div>
                <a href={`url/${fileinfo}`} target="_blank">fileinfo</a>
            </div>
        )
    }

    render() {
        return (
            <div>
                {this.state.files.map((val: any, idx: any, arr: any) => {
                        return this.generateFileData(val);
                })}
            </div>
        )
    }
}
