import {Select} from "antd";
import {CustomerServiceFilled} from "@ant-design/icons";
import React from "react";


const SelectOption = Select.Option;
const LayoutSelector = (props: { value: any; onChange: any; options: any; }) => {
    const {value, onChange, options} = props;
    return (
        <div>
            <Select style={{width: '120px'}} value={value} onChange={onChange}>
                {options.map((item: { type: any; icon: JSX.Element }) => {
                    const {type} = item;
                    const iconComponent = item.icon || <CustomerServiceFilled/>;
                    return (
                        <SelectOption key={type} value={type}>
                            {iconComponent} &nbsp;
                            {type}
                        </SelectOption>
                    );
                })}
            </Select>
        </div>
    );
};

export default LayoutSelector
