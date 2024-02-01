import {ProCard} from '@ant-design/pro-components';
import React from 'react';
import {Space, Typography} from "antd";
import {formatDate} from "@/utils/format";
import {Column} from "@ant-design/plots";

interface Props {
    hhi: Graph.HHILanguage
}

const {Text} = Typography

const HHITable: React.FC<Props> = (props) => {
    const hhi = props.hhi
    const paletteSemanticRed = '#F4664A'
    const brandColor = '#5B8FF9'
    // const getColor = (score: number) => {
    //     if (score > 60) {
    //         return paletteSemanticRed
    //     } else if (score > 30) {
    //         return brandColor
    //     }
    //     return 'green'
    // }
    const title = `${hhi.language} 软件垄断程度`
    const table = () => {

        const config = {
            data: hhi.hhIs,
            xField: 'name',
            yField: 'score',
            colorField: 'score',
            label: {
                text: (d: Graph.HHI) => `${(d.score).toFixed(1)}%`,
                textBaseline: 'bottom',
            },
        }
        return <Column {...config} />;
    }
    return <ProCard title={<Space direction={"horizontal"}>
        {title}
        <Text type="secondary">{'最后更新：' + formatDate(hhi.updateTime)}</Text>
    </Space>}>
        {table()}
    </ProCard>

};

export default HHITable;
