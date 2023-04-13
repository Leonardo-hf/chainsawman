import Guide from '@/components/Guide';
import { trim } from '@/utils/format';
import { PageContainer } from '@ant-design/pro-components';
import { useModel } from '@umijs/max';
import styles from './index.less';
import { request } from '@umijs/max';
import {connect} from "umi"

const HomePage: React.FC = (props) => {
  const { name } = useModel('global');
    console.log(mapStateToProps(props).graph)
    // return (
    //   <PageContainer ghost>
    //     <div className={styles.container}>
    //       <Guide name={trim(name)} />
    //         <button>wdnmd</button>
    //     </div>
    //   </PageContainer>
    // );
    return <div onClick={() => {
        props.dispatch({
            type: 'graph/queryGraphs'
        });

    }}>aaa

    </div>
};

const mapStateToProps = (state) =>{
    return {
        graph:state.graph,
    }
}

export default connect(mapStateToProps)(HomePage);
