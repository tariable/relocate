import React from 'react';
import IPage from '../../interfaces/IPage';
import Breadcrumbs from '../../common/Breadcrumbs/Breadcrumbs';
import FiltersBlock from './FiltersBlock/FiltersBlock';
import './AccommodationListPage.less';
import AccommodationList from '../../common/AccommodationList/AccommodationList';
import SortBlock from './SortBlock/SortBlock';

const AccommodationListPage: React.FC<IPage> = () => {
  const breadcrumbs = [{ text: 'Главная', link: '/' }, { text: 'Поиск' }];
  return (
    <div className="list-page">
      <div className="list-header">
        <Breadcrumbs breadcrumbs={breadcrumbs} />
        <SortBlock />
      </div>
      <div className="list-body">
        <FiltersBlock />
        <AccommodationList />
      </div>
    </div>
  );
};

export default AccommodationListPage;
