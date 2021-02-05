import React, {useState, Suspense} from 'react';
import SiteHome from './pages/home'
import './App.css';
import 'antd/dist/antd.css';
import {IntlProvider} from 'react-intl-hooks';
import {BrowserRouter, Switch, Route, Redirect} from "react-router-dom";
import {LoginUser} from "./model/user";
import {Locale} from "./components/LocaleSwitch/LocaleSwitch";
import {ConfigProvider, Result, Button, BackTop} from 'antd'
import zhCN from 'antd/lib/locale/zh_CN';
import {Locale as AntLocale} from 'antd/lib/locale-provider';
import routes, {RouteWithSubRoutes} from './pages/routes'
// import 'moment/locale/zh-cn';
import moment from "moment";

moment.locale('zh-cn');


interface LoginInfo {
    user: LoginUser | null
    onLangChange: (locale: Locale) => void
    login: React.Dispatch<React.SetStateAction<LoginUser | null>>
}

export const GlobalContext = React.createContext<LoginInfo | null>(null)

let defaultConfig = require('./locales/zh-Hans').default

/**
 * 获取国际化资源文件,
 * 其他语言懒加载
 */
async function getLocale(lang: string): Promise<any> {

    let result = []
    switch (lang) {
        case 'zh-CN':
            result[0] = defaultConfig
            result[1] = zhCN
            break;
        case 'en':
            result[0] = (await import(/* webpackChunkName: "en-US" */'./locales/en-US')).default
            result[1] = (await import('antd/lib/locale/en_US')).default;
            break;
        default:
            result[0] = defaultConfig
            result[1] = zhCN
    }

    return result
}


function App() {
    let [loginUser, setLoginUser] = useState<LoginUser | null>(null)
    const [langState, setLangState] = useState({
        lang: 'zh-CN',
        messages: defaultConfig.messages,
        formats: defaultConfig.formats
    });
    const [antLang, setAntLang] = useState<AntLocale>(zhCN)

    async function onLangChange(locale: Locale) {
        let [config, antLang] = await getLocale(locale.code)
        setLangState({lang: config.code, messages: config.messages, formats: config.formats})
        setAntLang(antLang)
        console.info('change lang success')
    }

    return (

        <BrowserRouter>
            <IntlProvider
                locale={langState.lang}
                messages={langState.messages}
                formats={langState.formats}
                defaultLocale="en"
            >
                <GlobalContext.Provider value={{user: loginUser, onLangChange, login: setLoginUser}}>
                    <ConfigProvider locale={antLang}>
                        <Suspense fallback={<div>Loading...</div>}>
                            <BackTop/>
                            <Switch>
                                <Route exact path="/">
                                    <SiteHome/>
                                </Route>

                                {routes.map((route, i) => (
                                    <RouteWithSubRoutes key={i} {...route} />
                                ))}


                            </Switch>

                        </Suspense>
                    </ConfigProvider>
                </GlobalContext.Provider>
            </IntlProvider>
        </BrowserRouter>
    );
}

export default App;
