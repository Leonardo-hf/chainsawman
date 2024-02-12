import {PageContainer} from "@ant-design/pro-components";
import React, {useEffect, useState} from "react";
import {connect} from '@@/exports';
import Markdown from 'react-markdown'
import remarkMath from "remark-math";
import remarkGfm from 'remark-gfm'
import rehypeKatex from 'rehype-katex'
import 'katex/dist/katex.min.css'
import styles from './juejin.css'
import {Anchor, Col, Row} from "antd";

type Title = {
    key: string;
    href: string;
    title: string;
    children?: Title[];
    nodeName: any;
}

const mapStateToProps = (state: any) => {
    return {
        docs: state.algo
    }
}

const AlgoDoc: React.FC<{ algoID: number, docs: any, dispatch: any }> = (props) => {
    const {algoID, docs, dispatch} = props
    /**
     * 格式化markdown标题的dom节点数组
     */
    const formatNavItem = (headerDom: NodeListOf<HTMLElement>) => {
        // 将NodeList转换为数组，并提取出需要的属性
        /**
         * (双重循环，从后往前，逐渐将子节点存入父节点children属性)
         * 1. 从后往前，将子标题直接存入前一个父级标题的children[]中
         * 2. 如果前一个标题与当前标题(或标题数组)无直系关系，则直接将当前标题(或标题数组解构后)放入list数组
         * 3. 循环多次，直到result数组长度无变化，结束循环
         */
        let result = Array.prototype.slice
            .call(headerDom)
            .map((item, index) => {
                return {
                    href: '#' + index,
                    key: '' + index,
                    title: headerDom[index].innerText,
                    children: [],
                    nodeName: item.nodeName,
                };
            }) as Title[];
        let preLength = 0;
        let newLength = result.length;
        let num = 0;
        while (preLength !== newLength) {
            num++;
            preLength = result.length; // 获取处理前result数组长度
            let list: Title[] = []; // list数组用于存储本次for循环结果
            let childList: Title[] = []; // childList存储遍历到的兄弟标题，用于找到父标题时赋值给父标题的children属性
            for (let index = result.length - 1; index >= 0; index--) {
                if (
                    // 当前节点与上一个节点是兄弟节点，将该节点存入childList数组
                    result[index - 1] &&
                    result[index - 1].nodeName.charAt(1) ===
                    result[index].nodeName.charAt(1)
                ) {
                    childList.unshift(result[index]);
                } else if (
                    // 当前节点是上一个节点的子节点，则将该节点存入childList数组，将childList数组赋值给上一节点的children属性，childList数组清空
                    result[index - 1] &&
                    result[index - 1].nodeName.charAt(1) <
                    result[index].nodeName.charAt(1)
                ) {
                    childList.unshift(result[index]);
                    result[index - 1].children = [
                        ...(result[index - 1].children as []),
                        ...childList,
                    ];
                    childList = [];
                } else {
                    // 当前节点与上一个节点无直系关系，或当前节点下标为0的情况
                    childList.unshift(result[index]);
                    if (childList.length > 0) {
                        list.unshift(...childList);
                    } else {
                        list.unshift(result[index]);
                    }
                    childList = [];
                }
            }
            result = list;
            newLength = result.length; // 获取处理后result数组长度
        }
        return result;
    };

    /**
     * markdown锚点注入方法
     */
    const getAnchors = () => {
        // 获取markdown标题的dom节点
        const header: NodeListOf<HTMLElement> = document.querySelectorAll(
            `.${styles.markdown} h1, .${styles.markdown} h2, .${styles.markdown} h3, .${styles.markdown}
            h4, .${styles.markdown} h5, .${styles.markdown} h6`
        )
        // 向标题中注入id，用于锚点跳转
        header.forEach((navItem, index) => {
            navItem.setAttribute("id", index.toString());
        });
        // 格式化标题数组，用于antd锚点组件自动生成锚点
        return formatNavItem(header)
    }

    /**
     * 锚点item点击事件
     * 1.解决antd的Anchor组件会在导航栏显示"#锚点id"的问题，
     * 2.以及本项目中navbar通过监听屏幕滚动进行定位，通过scrollIntoView设置页面滚动缓冲，可以一定程度上解决在页面快速滚动时navbar的定位切换造成的闪烁问题。
     * 当然也可以不设置该点击事件
     */
    const handleClickNavItem = (e: any, link: any) => {
        e.preventDefault();
        if (link.href) {
            // 找到锚点对应得的节点
            let element = document.getElementById(link.href);
            // 如果对应id的锚点存在，就跳滚动到锚点顶部
            element &&
            element.scrollIntoView({block: "start", behavior: "smooth"});
        }
    }
    const [sourceMd, setSourceMd] = useState("");
    const [titles, setTitles] = useState<Title[]>([])
    useEffect(() => {
        if (algoID in docs) {
            setSourceMd(docs[algoID])
            return
        }
        dispatch({
            type: 'algo/queryAlgoDoc',
            algoID: algoID
        })
    }, [docs, algoID])

    useEffect(() => {
        setTitles(getAnchors())
    }, [sourceMd])

    return <PageContainer>
        <Row style={{position: 'relative'}} gutter={12}>
            <Col span={18}>
                <Markdown className={styles.markdown} remarkPlugins={[remarkGfm, remarkMath]}
                    // @ts-ignore
                          rehypePlugins={[rehypeKatex]}>{sourceMd}</Markdown>
            </Col>
            <Col span={6}>
                <aside
                    className={`${styles.aside} container`}
                    style={{position: "fixed", top: 72}}
                >
                    {titles.length > 0 && (
                        <Anchor
                            affix={false}
                            offsetTop={100} // 设置距离页面顶部的偏移
                            onClick={handleClickNavItem}
                            items={titles}
                        />
                    )}
                </aside>
            </Col>
        </Row>
    </PageContainer>
}

export default connect(mapStateToProps, null)(AlgoDoc)