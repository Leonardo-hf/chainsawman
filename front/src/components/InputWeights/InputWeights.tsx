import {ProColumns, ProFormItemProps, ProFormSlider, ProTable} from '@ant-design/pro-components';
import React, {useState} from 'react';

interface Props {
    innerProps?: ProFormItemProps,
    max?: number,
    step?: number,
    headers: string[]
}

export const getWeights = (w: number[]) => {
    const weights: string[] = []
    for (let i = 1; i < w.length; i++) {
        weights.push((w[i] - w[i - 1]).toString())
    }
    return weights
}

const InputWeights: React.FC<Props> = (props) => {
    const max = props.max ? props.max : 1
    const min = 0
    const step = props.step ? props.step : 0.0001
    const interval = (max - min) / props.headers.length
    const initValues = [min]
    for (let i = 0, last = min; i < props.headers.length; i++) {
        last += interval
        initValues.push(last)
    }
    const fixed = step >= 1 ? 0 : Math.ceil(Math.log10(1 / step))
    const values2Rows = (v: number[]) => {
        const vmap: any = {'sum': max}
        for (let i = 1; i < v.length; i++) {
            vmap[props.headers[i - 1]] = (v[i] - v[i - 1]).toFixed(fixed)
        }
        return [vmap]
    }
    const [rows, setRows] = useState<any[]>(values2Rows(initValues))
    const columns: ProColumns[] = props.headers.map(h => {
        return {
            title: h,
            dataIndex: h,
            fixed: 'left',
        }
    })
    columns.push({
        title: '合计',
        dataIndex: 'sum',
        fixed: 'left'
    })
    return <>
        <ProFormSlider
            initialValue={initValues}
            {...props.innerProps}
            range
            fieldProps={{
                tooltip: {
                    formatter: (v) => {
                        return (v! * 100 / (max - min)).toFixed(2) + '%'
                    }
                },
                // @ts-ignore
                defaultValue: initValues,
                onChange: (v: number[]) => {
                    setRows(values2Rows(v))
                }
            }}
            step={step}
            max={max}
        />
        <ProTable
            rowKey={(row)=>row}
            cardProps={{bodyStyle: {padding: 0}}}
            dataSource={rows}
            columns={columns}
            search={false}
            options={false}
            pagination={false}/>
    </>

};

export default InputWeights;
