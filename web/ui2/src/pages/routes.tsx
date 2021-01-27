import SiteLayout from "../layouts/SiteLayout";
import React from "react";
import {Route} from "react-router-dom";

export function RouteWithSubRoutes(route: any) {
    return (
        <Route
            exact={route.exact}
            path={route.path}
            render={props => (
                // pass the sub-routes down to keep nesting
                <route.component {...props} routes={route.routes}/>
            )}
        />
    );
}
const routes = [

    {
        path: "/",
        component: SiteLayout,
        routes: [
            {
                exact:true,
                path: "/topic/:topicId",
                component: React.lazy(() => import('./topic'))
            },
            {
                path: "/account",
                component: React.lazy(() => import('./account')),
            }, {
                exact:true,
                path: "/space",
                component: React.lazy(() => import('./space')),
            },{
                exact:true,
                path: "/login",
                component: React.lazy(() => import('./account/login'))
            }
        ]
    }

];

export default routes
