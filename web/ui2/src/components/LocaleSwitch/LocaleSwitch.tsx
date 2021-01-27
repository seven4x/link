import React, {useState} from 'react'
//https://stackoverflow.com/questions/61015828/react-svg-tag-name-provided-is-not-valid
import {ReactComponent as Check} from './check.svg'

import "./localswitch.css"
import useUrlState from '@ahooksjs/use-url-state';

export interface LocaleSwitchProps {
    defaultLocale?: string
    onLocaleChange: (locale: Locale) => void
}

export interface Locale {
    code: string
    language: string
    flag: string
}

const locales = [
    {
        code: 'zh-CN',
        language: '简体中文',
        flag: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADwAAAA8CAMAAAANIilAAAAAOVBMVEUAAAD/PAD/PAD/PQD/PQD/6zv/SAP/Uwf/wCz/kx3/qCT/4Df/cxL/aA7/1TP/tSj/fxb/XQv/iRpQIAmlAAAABHRSTlMAv4BgC6AmpAAAALNJREFUSMft1rsOwyAMBdA2qR/lESD//7EttAphqGTcJYPvwoCODFcM3CyWS2e9q7I2vDxUWX5h/Adn1OO4MaoxJjc/OR9XxmmMMBicwgHC0BrNYAdubI1kOD7f8eDrEr+n5iSdvHv4xO+9NfGxiZtl0j3PhpVvO1R77jugHDtIlE59R3YkxluuvW1917O8sNKWckwjLCIsiRwH1ODelQL3rpS4d6UtzLBhw3qs/9BYLFfOC1AUIa4gRIvdAAAAAElFTkSuQmCC'
    }, {
        code: 'en',
        language: 'English',
        flag: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADwAAAA8CAMAAAANIilAAAAAhFBMVEUAAADvjYbwfHSKYpnt7/GKYpp4hcrr7/HwkYnvq6WAd7Xvx8aHaqShaJDLsMOic5vykYzyj4r0Qzbs7/E/UbXvmJPybmTuxMJ7h8zBp77///+HktFvfMfb3/FLW7nz9Pu3vuNjccNsTZSrs99qeMSTndXn6fbP1O3O0+xrXKbNrr9rbbgEiDpcAAAAEnRSTlMAcMDgfuDAgbCgzI7Y2J6UPDtbvABhAAAAz0lEQVRIx+2TSxKCMBBEBQEB/zBgTEIUERS9//1UilUkMRW0KKy8Re/epqdnYjD8L+upiKUtwWpkF0ScIgn2V2QCwAelSnKYZ4cjcFFWVQn3nYR5I3uQYQRcMIwZnGMJTisjRIELRoiiLEBFdqGzMVYqtn3BV8YFfdamJhe44KKpTUFevI5K3oKgvE4l+P0L8wQLU5NFC1ORQ+HC6kRC0LYtWFivl/y5vNqLuH1+yVmshTOwbEda2APLfqqFP+ZTBYkWwZjvvLG02E4MBkMXD5MAkJrGRHOLAAAAAElFTkSuQmCC'
    }, {
        code: 'fr-FR',
        language: 'Français',
        flag: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADwAAAA8BAMAAADI0sRBAAAAIVBMVEUAAADs7/A+ULX/PAA/ULT/PAA9ULX/PQDs7/E/UbX/PQAILXrPAAAACHRSTlMAv4CAv79gYG3QJLwAAAAySURBVDjLYxgFtAApLlAgCAehUBAOlNacCQUdcLAKChaPSo9Kj0qPSpMlTajoGQU0AADZmVPJRnX9jgAAAABJRU5ErkJggg=='
    }, {
        code: 'de-DE',
        language: 'Deutsch',
        flag: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADwAAAA8BAMAAADI0sRBAAAALVBMVEUAAABTV1z/mAVEWWT/wAf/PABEWmT/wQZlVFH/nARFWmL/wgX/PQBFWmT/wQcERp1ZAAAADHRSTlMAz8+/v4CAgHR0YGBk1e17AAAARklEQVQ4y2MYBbQAi41xAiugdO5dnODaqDSZ0o2COIEEUDrmDE5wdFSaNtKTlHACTaB03Tuc4PmoNJnSW1xwAm+GUUADAAA0cjjQYLRUuAAAAABJRU5ErkJggg=='
    }, {
        code: 'it-IT',
        language: 'Italiano',
        flag: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADwAAAA8BAMAAADI0sRBAAAAIVBMVEUAAADs7/Bonzj/PABonzj/PABonzj/PQDs7/Fonzj/PQCUdpRIAAAACHRSTlMAv4CAv79gYG3QJLwAAAAySURBVDjLYxgFtAApLlAgCAehUBAOlNacCQUdcLAKChaPSo9Kj0qPSpMlTajoGQU0AADZmVPJRnX9jgAAAABJRU5ErkJggg=='
    }, {
        code: 'ja-JP',
        language: '日本語',
        flag: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADwAAAA8BAMAAADI0sRBAAAAJ1BMVEUAAADs7/Dr7/Hs7/Ls7/HVAADns7TWDg/r4OHnwsPjlZbeWVrbOzyRGGwAAAAABHRSTlMAv4BgC6AmpAAAAJBJREFUOMtjGAW0AIaCOIEwUFrFBSdwGkLSbRl4pN22h4ZWp+CS9jwaCgQxU3BILw8Fgyrs0l6hULAEq/R0mHQlVumjMOkYbNJuoXCQgkXaGyG9BYu0O0K6BIt0K0I6Aot0KkI6jIA0YcMJO42wxwgHC2agkhMliAglIzkgEhPJSRGRkAdtHiNNmlDRMwpoAADdAyGYeDi9yQAAAABJRU5ErkJggg=='
    }, {
        code: 'pt-PT',
        language: 'Português',
        flag: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADwAAAA8CAMAAAANIilAAAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAAHIUExURUdwTP/fAP/wAP/pAP/gAACdKwAieQCcLQCcLQCcLf/sAACZMf//AAAfdwCQOwCSOQCWMwCUNgCMP//iAACOPf/rAAAbdQCNPwAlfP/yAP/uAAAJhv/eAAAWcgCjIgCbLf/lAP/kAP/3AAAYfQAkef/nAACzAP/9AAA/iQAsfgCgJgA5hwAAmAA0g//1AAAZdP/6AAJDjQAAZgARcAADcgCoGtPZAACsEfLhAAAwgQAeewATgQAAkwADiAABj8nWAACaMFy+AKvQAPrjAGZyReDAAENaUQABn+z/6Dm6AJXKALfTAHfDAACvCp7MAIvIANzcAAAJa8avAAAvcW97QlxpSwAAbBc8jWm6rgAKegAKcr/VAEy8AIPFAGrBAAk3beffANzKAKGWFJ/kwzFQVrXwz0xxoSJOXgAWlmybvYjPuQAAXz1dmZDTvou/ux63ACi3ADlOXBu2AHmntGmorNXEAMSvAP///cu9AAALjQA8e7KhADdMXQAoilF8qh5FTjhtnSc5WrCfBniBOunOAI+JGMu2MP/qOwAorGasr47bv6zl74l6AG+kp///owBCqUyDnWyIe3/F6gAeVaP4xLf90JD159L/2s6Vw/IAAAAJdFJOUwD////////y5X/OHOMAAANxSURBVEjH7ZZXd9pKEIAjtCvHK4lmIYkFgWRYuulgAwbbcU3ixL0muem93N57eu/l70YSxDbEiQk5OScPnhce9nyMdr+dmd2zZzd242uOvV0dxV4T7u4wTLirM7ZrBxja7bBDGPqFWEzww05g6BAODQyEBBv8ZBgOO/0ziKbRAehk4afB0GYNDeAeiurBoz9Yt03+IRiyTt8+TEgyHE4Sgg/6nNsc3PYw7LY6jqqol1PH5+fHNa4XqSMOoRu2A0O7szCELNzJseMBjydwfOwkZ0HFguSAO8LQL9mnaeJO3/AMegLlckD/WexNEtc0KzVbex82/MxiYElOJCpBxoxgJfG92wLwiRZrrTBkJXgA0ZTKTSR4ntmIxESYUDQ6DKUt1lrgd356wg9r+fV8PpudSzGM/iflxCnOa1j71raZvAlu+KEAAOjRxdXV1XsX/8uUl7Oi/uXVBQ0BQDVZ2wLDbsFxlGBgsNSdB/evXrp0emrq19M//RO/nAsyteucsQAQ2bC2CUOHVBjCqoECkF68u5JfWwswK+d+PDt19mkgmxpcCptLlIqL/XVrmzAbm6YRqLNa8t9aVOFTOWUuuxy4cu3Ny3P5tfNYBXXc5ToWY7fCrH0orHrrq0BF56siw0fjfFAUlfX8yrWpqxde0KSx7FXDRZZtyuybQa56YqC5z9QUXo728alSSWH47PKVV6+febVGZr3QfGzLnoX+4rs9c2PVeEZUeF6OREoiE2Qurz9/7DXXKA3PHhJa9my0DftIY9vcnxcyemKR72PEuG5KjIiTt9P106a/sTeaS7Nnu9N3xPTs0v6o5jIlRc5FInM8I+YyT5SbFvOohwrbea6Xom2/pqv2cqcGy3w8o6Ti8aDMR6OZyVv6DQNY3W+zbpTm+3fbCQ8jWj/PiYQYLfFySpHjfZHI5Jk08RK8L+b88N02r7cQOoGB2/JXosIHGVlW+qJ/T/6PkxQeCFk/WlVm45PYYy6S7F2s6vUsywFP7fcxd1JFM35pGO7cSRxSf9Fl4X5eWqh4PJWFpd84t2u2v51O0qiRERWlOWp8/pdxL5dGZMNPO91Tt3YQE9ronjTBRwrtd0+Dhlbbd6NG327x0+bEMKxhQlr8tD2rdGujWsjawaxqTEmfMAw7ns8O2PFwb+Nl8Flvks96De3GbnzBeAtrG6wXpnakqQAAAABJRU5ErkJggg=='
    }, {
        code: 'ru-RU',
        language: 'Русский',
        flag: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADwAAAA8BAMAAADI0sRBAAAALVBMVEUAAADU1d7DQzfg4OL/PAA+ULXf4eH/PADEyNrKQjPf4uL/PQA/UbXg4eL/PQCA2eEDAAAADHRSTlMAz8+/v4CAgHR0YGBk1e17AAAARklEQVQ4y2MYBbQAi41xAiugdO5dnODaqDSZ0o2COIEEUDrmDE5wdFSaNtKTlHACTaB03Tuc4PmoNJnSW1xwAm+GUUADAAA0cjjQYLRUuAAAAABJRU5ErkJggg=='
    }, {
        code: 'es-ES',
        language: 'Español',
        flag: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADwAAAA8BAMAAADI0sRBAAAAG1BMVEUAAADcLADdLADzjwTvfgPcKwDdLAD/wQf2mwV9JkN3AAAABnRSTlMAv4DAsGApsHfPAAAAOUlEQVQ4y2MYBbQAgYI4gShQWi0NJ0gawdIeHThBC1DavBwnKB6VHn7S+JLD4E7IAyZNqOgZBTQAAMrqsZDrvRzeAAAAAElFTkSuQmCC'
    }, {
        code: 'ko-KR',
        language: '한국어',
        flag: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADwAAAA8CAYAAAA6/NlyAAAABmJLR0QA/wD/AP+gvaeTAAAINElEQVRoge2Zd1RUZxqHnylShEFAgzGuBRQ7okCwrEZXo1nXcjRoLKAC4oINQxmFRBSxBKQJCliIZcWGrjmxHSXGEDeiogaJu8bIrkalF1GRJjJ3/3A1a2Zw7ox4TM6Z568593u/9/v95t5vvve9AwYMGDBgwIABAwYMGNAHybMPVVVVw6VS6fA3J+X1oVKpMhUKRSaA/NnF/5ld8WYkvV6kUilAJoD0jSp5A+htWKVScSbrfHNqEcXlK7k0NjbqPV9vw3sOHGLXvnQEQaCmplZvAWIpKimlsbGRU5ln2LXvgN559DJcWFTMtrQ9hC0JIuN0JpHxiXoLEEtMYhK79h3A39eHIycyuPnzbb3y6GxYEARWrYvDy306rSwUJH++g/lzvbiUk0t80ma9RLyM+KTNXL6SS7D/AnbuTaeopJSQgEVERMXQqFLpnE9nw4eOHKO6poYZH7mxLiGJmVOnYG1pRdT6DYwZPVJnAdqYNP4vrI6Op5VCgZ/3LCKiYnDu54irsxMlJaU659PJcFl5BSmpO1gREsz57EuUlpXhNmEsyanbGDlsKD3su+osQBudO3ZgzKiRbNq+kykTJyCXy0n/4jDz5niSuDmVqqpHOuWTaw/5hbVxCUx1m0i7t9sS9OkKEqPWcOPfNzl74SJpW5NfPvl+JWSfRcj7CSrvPb1m3RpJ1+7gOhgsrZqc6u0xnZl/XcD1G3mEKQPxnL+YIYNc6efQh9iNKYSHKkV7EG34eMYpCgqLiI5YTnRiMpPGj6X9O+2Y6buQZcEBGBsZaZ5YXoqQtg2++QoE9T0ncAQkUhgxGom7F7SxURcplxMfuYqi4hJ6du+G54ypRETFkRIXhe/HwWRduMjgAe+K8iHqka68/4C4pM2sDFXyr+s3uHrtR9ynfMi2tD249HfE0aG35ok5FxEWzYHTJzWa/cW1Cr4+gbDAEy5kaQx528aGQ0eOcfDLo3hMm0JNbQ1HT35FmDKQdQkbRR+NogxbKMyJDF9GF1tbVkfHsXxpEAVFxRzP+Bo/r9maPWSdQVgZAtU67LHaWoTPwhCyzmgcVvovIC39IGXl5YSHKElI2YKpqSnRq1ZgamoiaglRhmUyGS79HamuqcHbYzrdutixMjKGpQGLaNnSVH3C7VuwPhL0ODZQqZ7OvXNLbchCoWCBjxdrotfT1c6WoIXzaFSpsO9ih0Qi0ZBMHZ1+tKwsWzFm1EjyCwvp37cPA12cNcYJWzdAnbhH7J5lB451GkdegyUA9i3uM/b2Uay3bESyOlYtfvSI4VzKyaWsvIKxH4zSRT7wf+1hdXV1OM3QLQnXrkKIv6jY893HsrbIgdq6F2tjUxMZn7S7ykCvD5D0cnhVSQArzczMwuE1dEuSrG9Fxf38Tj9W5fdRMwtQW9fI6vw+3PlBv/LxZTS7YeHqFVFxu63/zOOGpvd4fYOKtJtNn8360vz9cHmZqLCcIu1L5/yn+buw5jdcX6c15LGRKQ8fNWiNq6nVv+9timY3XN3OTmvMXZteCIL2XK2tjJtB0Ys0u+FrvUdrjTnR+j1RuTq0N3tVOWo0u+EKu/7kdHm/yfHvu47iaF4TdfevGOTStrlkPUcvw/WPH3Mnv0Dj2JDB7fns4UAOOPpQa6J4fr3WREF6Xx+WFTnx5In2CsyspZzhg9tpHCsqLtFHNqBjpQVQW1dHYVExS1esIm1rMibGL+4zhXkLJk+wZcuuena2mE9H2xYA3CltoP66ChBXbrq7dUVh3kLtenFpKX6BS9j7+SaMjYyQyWQ66dfpDpeWleM20xubNm0Y8d5QtmzfpTFuynhbBjjbUN+gIq+gnryCeupfcub+GmfHNriNs9U4Fhm/gUVzvVGpVEyeNYeKykpdLIgz/LCqivDIaFpbWzHsj4OIS97E3NnunL2QzY8/3VBPKpXwyeJ+9O1lrZMYgL69rFke5IRMpt4MHM84BYLA+38aRuzGFFydnVCYmROzIRlBzM8+ottDBaVl5aTtP4i/rw+Xvs/lcu4PLFMGEBEVy5MnT9TmmLWUE7XclcnjbZHLtS8jl0uZMt6OdctdMWupeacdyzjFMmUAl3KukH05h8Xz5pL6tzSMjYxEd0uiH+kwZSA79uynuLSMkIBFJKZsxaFXT5z6ObJzb7rGOS3kUvxm9yQ1figTx3TGpo16K2nTxpSJYzqTGj8U39k9mvxy1sTEExroj7WVFWtjE/g06GMKCos4feY75nrOFGtDt25pz4G/880/zrIlIZZbt+/y4OEDenSzx91nHjGrw7Hr3OmliwkCVN6vp6KyHnhaWFhZGqPt5hw5kUHa/oPs3prMt2fP0bO7PW3fegvPef4ELPDFybGvNp/6dUvT3CbR2Kji4JdHMTU1ITwyhoaGBkID/YmIikGlpeGXSMDayhh7Owvs7SywttJutuLePdYnbyY8VMmVq/8kLmkTlhatuFtQgM9sDzFmX0Anw1KplBUhwWzatgOJBKa7TSJu4ybedepP544dSf/isE6LiyEyfgNuE8bRuWMHIqJiWRYcQH5hIcqwCIYMGqBzPp0Lj04d/sA0t4msiV7PR5MmkF9YyLnsSwQu9MOlv6POArTx4fix+MzyIDl1O44OvXF1cSJiXSyhgf7IpLrXTXpVWt4eMyivqODk6UzClgRx+24+FgoFXe00n52vwiBXF+RyGVVVjwhaOI/d6Qfp1b27zo/yM/QyLJPJCFsSyKHDx57f8deJVColPFSJwtyc785l4+/no3euV3qnpVKpnv27/luned5p/U7MvsDvT7EBAwYMGDBgwICB3yj/BXEp1EvWHHaCAAAAAElFTkSuQmCC'
    }]

// 参考 https://icons8.cn/
const LocaleSwitch: React.FC<LocaleSwitchProps> = (props) => {
    let {defaultLocale, onLocaleChange} = props
    let defVal = locales[0]
    for (let l of locales) {
        if (l.code === defaultLocale) {
            defVal = l
        }
    }
    let [current, setCurrent] = useState(defVal)
    const [lang, setLang] = useUrlState<any>({lang: defVal.code})

    let [popupShow, setPopupShow] = useState<boolean>(false)
    const onChange = (selected: any) => {
        if (selected.code === current.code) {
            console.log('not change')
        }
        setLang({lang: selected.code})
        setCurrent(selected)
        setPopupShow(false)
        onLocaleChange(selected)
    }
    if (lang.lang !== current.code) {
        //reset from url lang code
        for (let l of locales) {
            if (l.code === lang.lang) {
                console.log("reset from url lang code")
                setTimeout(() => {
                    onChange(l)
                }, 50)
                break
            }
        }
    }

    return (
        <div className="popup menu-language header-language">
            <div className="popup_target">
                <div className="language-target" key={current.code + "_checked"} onClick={() => {
                    setPopupShow(!popupShow)
                }}>
                    <img src={current.flag} alt={current.code}/> <span>{current.language}</span>
                </div>
            </div>
            <div className="popup_content is-bottom-center animate-off" style={{display: popupShow ? '' : 'none'}}>
                <div className="languages has-flags">
                    {
                        locales.map((locale) => {
                            if (locale.code === current.code) {
                                return <div className="ls-is-active language" data-code={locale.code} key={locale.code}
                                            onClick={(e) => onChange(locale)}>
                                    <img src={locale.flag} alt={locale.code}/> <span>{locale.language}</span>
                                    <span className="icon-check">
                                                <Check width="16px" height="16px"/>
                                            </span>
                                </div>
                            } else {
                                return <div className="language" data-code={locale.code} key={locale.code}
                                            onClick={(e) => onChange(locale)}>
                                    <img src={locale.flag} alt={locale.code}/> <span>{locale.language}</span>
                                </div>
                            }
                        })
                    }
                </div>
            </div>
        </div>
    )
}

export default React.memo(LocaleSwitch)