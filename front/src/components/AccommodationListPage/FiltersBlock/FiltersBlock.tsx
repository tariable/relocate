import React from 'react';
import './FiltersBlock.less';
import MetroSelect from '../../../common/MetroSelect/MetroSelect';
import RoomSelect from '../../../common/RoomSelect/RoomSelect';
import PriceSelect from '../../../common/PriceSelect/PriceSelect';

const FiltersBlock: React.FC = () => (
  <div className="filter-block">
    <div className="filter-room-select">
      <RoomSelect />
    </div>
    <MetroSelect />
    <div className="filter-price-select">
      <PriceSelect />
    </div>
  </div>
);

export default FiltersBlock;
