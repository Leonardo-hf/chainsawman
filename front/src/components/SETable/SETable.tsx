import {ProCard, ProList} from '@ant-design/pro-components';
import React from 'react';
import {Space, Tag, Typography} from "antd";
import {formatDate} from "@/utils/format";
import {createFromIconfontCN} from "@ant-design/icons";

interface Props {
    hot: Graph.HotSETopic
}

const {Text} = Typography

const SETable: React.FC<Props> = (props) => {
    const hot = props.hot
    const title = `最近热门 ${hot.language} 软件`
    const IconFont = createFromIconfontCN({
        scriptUrl: [
            '////at.alicdn.com/t/c/font_4378832_9tnr6hpdx5u.js'
        ],
    });
    return <ProCard title={<Space direction={"horizontal"}>
        {title}
        <Tag color={'green'}>{hot.topic}</Tag>
        <Text type="secondary">{'最后更新：' + formatDate(hot.updateTime)}</Text>
    </Space>}>
        <ProList<Graph.HotSE>
            rowKey='homepage'
            dataSource={hot.software}
            metas={
                {
                    title: {
                        render: (_, r) => <a href={r.homePage} target={'_blank'}>{r.artifact}</a>
                    },
                    avatar: {
                        render: (_, r, i) => {
                            let p = 'icon-pypi'
                            if (r.homePage.startsWith('https://github.com')) {
                                p = 'icon-github-fill'
                            } else if (r.homePage.startsWith('https://gitee.com')) {
                                p = 'icon-gitee'
                            }
                            return <Space>
                                <>{i + 1}</>
                                <IconFont type={p} style={{ fontSize: '32px' }}/>
                            </Space>
                        }
                    },
                    actions: {
                        render: (_, r) => r.score.toFixed(4)
                    },
                    subTitle: {
                        dataIndex: 'version'
                    },
                }
            }
        />
    </ProCard>

};

export default SETable;
