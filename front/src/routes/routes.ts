import paths from './paths';
import IRoute from '../interfaces/IRoute';

import MainPage from '../pages/MainPage/MainPage';
import AccommodationList from '../pages/AccommodationList/AccommodationList';
import AccommodationView from '../pages/AccommodationView/AccommodationView';

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
