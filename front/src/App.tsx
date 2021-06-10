/* eslint-disable react/jsx-props-no-spreading */
import React from 'react';
import './App.less';
import {
  Switch, Route, RouteComponentProps, BrowserRouter,
} from 'react-router-dom';
import Routes from './routes/routes';
import Header from './common/Header/Header';
import Footer from './common/Footer/Footer';

const App: React.FC = () => (
  <div>
    <Header />
    <BrowserRouter>
      <Switch>
        {Routes.map((route) => (
          <Route
            key={route.name}
            path={route.path}
            exact={route.exact}
            render={(props: RouteComponentProps<any>) => (
              <route.component {...props} {...route.props} />
            )}
          />
        ))}
        <Route component={() => <div>not found</div>} />
      </Switch>
    </BrowserRouter>
    <Footer />
  </div>
);

export default App;
