import "react"
import "./ConfigView.css"
import {useEffect, useState} from "react"
import axios from "axios"

export const ConfigView = () => {
    const [config, setConfig] = useState({})
    const [isLoaded, setIsLoaded] = useState(false)

    const [sumDelay, setSumDelay] = useState(0)
    const [diffDelay, setDiffDelay] = useState(0)
    const [multiplyDelay, setMultiplyDelay] = useState(0)
    const [devideDelay, setDevideDelay] = useState(0)
    const NanoSecToSec = 1000000000

    const fetchConfig = () => {
        const url = "http://localhost:8000/get_config"
        axios.get(url).then((response) => {
            setConfig(response.data)
            setSumDelay(response.data.sumDelay / NanoSecToSec)
            setDiffDelay(response.data.diffDelay / NanoSecToSec)
            setMultiplyDelay(response.data.multiplyDelay / NanoSecToSec)
            setDevideDelay(response.data.devideDelay / NanoSecToSec)
        })
    }

    const onChangeSumDelay = (event) => {
        setSumDelay(event.target.value)
    }

    const onChangeDiffDelay = (event) => {
        setDiffDelay(event.target.value)
    }

    const onChangeMultiplyDelay = (event) => {
        setMultiplyDelay(event.target.value)
    }

    const onChangeDevideDelay = (event) => {
        setDevideDelay(event.target.value)
    }

    useEffect(() => {
        if (isLoaded) return

        fetchConfig()
        setIsLoaded(true)
    }, [config])

    const sendNewConfigSettings = () => {
        let data = {
            sumDelay: sumDelay * NanoSecToSec,
            diffDelay: diffDelay * NanoSecToSec,
            multiplyDelay: multiplyDelay * NanoSecToSec,
            devideDelay: devideDelay * NanoSecToSec,
        }

        const url = "http://localhost:8000/set_config"
        axios.post(url, data).then((response) => {
            alert("Данные отправились")
        })
    }

    return (
        <>
            <div className="config_settings">
                <div className="settings__title">
                    <h2>Config Settings</h2>
                    <div className="settings_buttons">
                        <button
                            className="update"
                            onClick={() => {
                                fetchConfig()
                            }}
                        >
                            Update Config
                        </button>

                        <button
                            className="set_new_value"
                            onClick={() => {
                                sendNewConfigSettings()
                            }}
                        >
                            Set New Value
                        </button>
                    </div>
                </div>

                <div className="settings">
                    <p>Config values in seconds</p>
                    <div className="settings_values">
                        <div className="settings__operation">
                            <p>Sum Delay</p>
                            <input value={sumDelay} onChange={onChangeSumDelay}></input>
                        </div>
                        <div className="settings__operation">
                            <p>Diff Delay</p>
                            <input value={diffDelay} onChange={onChangeDiffDelay}></input>
                        </div>
                        <div className="settings__operation">
                            <p>Multiply Delay</p>
                            <input value={multiplyDelay} onChange={onChangeMultiplyDelay}></input>
                        </div>
                        <div className="settings__operation">
                            <p>Devide Delay</p>
                            <input value={devideDelay} onChange={onChangeDevideDelay}></input>
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}
