import React from "react";
import {render} from "react-dom";
import App from "./containers/App";
import EventCreate from "./components/Event/Create";
import EventList from "./components/Event/List";
import GuestCode from "./components/Guest/Code";
import GuestCreate from "./components/Guest/Create";
import {Router, Route, IndexRoute, browserHistory} from "react-router";

render(
    <Router history={browserHistory}>
        <Route path='/' component={App}>
            <IndexRoute component={GuestCode} />
            {/*unregistered route*/}
            <Route path='guest/code' component={GuestCode} />
            {/*admin routes*/}
            <Route path='admin/event/create' component={EventCreate} />
            <Route path='admin/event/list' component={EventList} />
            <Route path='admin/guest/create' component={GuestCreate} />
        </Route>
    </Router>,
    document.getElementById('root')
);