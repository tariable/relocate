import React from 'react';
import './FiltersBlock.less';
import MetroSelect from '../../../common/Filters/MetroSelect/MetroSelect';
import RoomSelect from '../../../common/Filters/RoomSelect/RoomSelect';
import PriceSelect from '../../../common/Filters/PriceSelect/PriceSelect';
import MetroDistance from '../../../common/Filters/MetroDistance/MetroDistance';

const FiltersBlock: React.FC = () => (
  <div className="filter-block">
    <div className="filter-room-select">
      <RoomSelect />
    </div>
    <div className="filter-metro-select metro-select">
      <MetroSelect mode="multiple" />
    </div>
    <div className="filter-metro-distance">
      <MetroDistance />
    </div>
    <div className="filter-price-select">
      <PriceSelect />
    </div>
  </div>
);

export default FiltersBlock;
