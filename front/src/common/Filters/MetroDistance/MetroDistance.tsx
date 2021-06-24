import React from 'react';
import './MetroDistance.less';
import { Select } from 'antd';

const { Option } = Select;

const distanceTypeArr = [
  {
    name: 'Неважно',
  },
  {
    name: 'Пешком',
  },
  {
    name: 'Транспортом',
  },
];

const MetroDistance: React.FC = () => (
  <div>
    <div className="metro-distance-title">
      Расстояние до метро
    </div>
    <Select
      size="large"
      defaultValue={distanceTypeArr[0].name}
    >
      {distanceTypeArr.map((type) => (
        <Option value={type.name} key={type.name}>
          {type.name}
        </Option>
      ))}
    </Select>
  </div>
);

export default MetroDistance;
