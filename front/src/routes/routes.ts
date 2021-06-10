import { lazy } from 'react';
import paths from './paths';
import IRoute from '../interfaces/IRoute';

const MainPage = lazy(() => import('../pages/MainPage/MainPage'));
const AccommodationList = lazy(() => import('../pages/AccommodationList/AccommodationList'));
const AccommodationView = lazy(() => import('../pages/AccommodationView/AccommodationView'));

const routes: IRoute[] = [
  {
    path: paths.mainPage,
    exact: true,
    name: 'Main page',
    component: MainPage,
  },
  {
    path: paths.accommodationList,
    exact: true,
    name: 'Accommodation List',
    component: AccommodationList,
  },
  {
    path: paths.accommodationList,
    exact: false,
    name: 'Accommodation View',
    component: AccommodationView,
  },
];

export default routes;
