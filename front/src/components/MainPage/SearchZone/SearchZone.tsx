import React from 'react';
import './SearchZone.less';
import MetroSelect from '../../../common/MetroSelect/MetroSelect';
import RoomSelect from '../../../common/RoomSelect/RoomSelect';
import PriceSelect from '../../../common/PriceSelect/PriceSelect';

const SearchZone: React.FC = () => (
  <div className="search-zone">
    <div className="title">
      Найди квартиру мечты в Москве с Relocate
    </div>
    <div className="search-body">
      <MetroSelect />
      <RoomSelect />
      <PriceSelect />
    </div>
  </div>
);

export default SearchZone;
