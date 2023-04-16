import Guide from '@/components/Guide';
import {trim} from '@/utils/format';
import {PageContainer} from '@ant-design/pro-components';
import {useModel} from '@umijs/max';
import styles from './index.less';

import {connect} from "umi"

const HomePage: React.FC = (props) => {
    const {name} = useModel('global');
    console.log(mapStateToProps(props).graph)
    return (
        <PageContainer ghost>
            <div className={styles.container}>
                <Guide name={trim(name)}/>
            </div>
            <div onClick={() => {
                props.dispatch({
                    type: 'graph/queryGraphs'
                });
            }}>aaa
            </div>
        </PageContainer>
    );
};

const mapStateToProps = (state) => {
    return {
        graph: state.graph,
    }
}

export default connect(mapStateToProps)(HomePage);
