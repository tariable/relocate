import React from 'react';
import IPage from '../../interfaces/IPage';
import AccommodationList from '../../common/AccommodationList/AccommodationList';

const AccommodationListPage: React.FC<IPage> = () => (
  <div>
    Accommodation list page
    <AccommodationList />
  </div>
);

export default AccommodationListPage;
