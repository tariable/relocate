import React from 'react';
import IPage from '../../interfaces/IPage';
import BenefitList from './BenefitList/BenefitList';
import './MainPage.less';
import AccommodationSuggestions from './AccommodationSuggestions/AccommodationSuggestions';

const MainPage: React.FC<IPage> = () => (
  <div className="main">
    <BenefitList />
    <AccommodationSuggestions />
  </div>
);

export default MainPage;
