import Guide from '@/components/Guide';
import { trim } from '@/utils/format';
import { PageContainer } from '@ant-design/pro-components';
import { useModel } from '@umijs/max';
import styles from './index.less';
import {useState} from "react";
import {connect} from "umi"
import { request } from '@umijs/max';

const HomePage: React.FC = (props) => {
  const { name } = useModel('global');

  console.log(props.graph.graphs)
  // return (
  //   <PageContainer ghost>
  //     <div className={styles.container}>
  //       <Guide name={trim(name)} />
  //         <button>wdnmd</button>
  //     </div>
  //   </PageContainer>
  // );
    return <div onClick={() => {
        request('/api/graph/getAll')
        // props.dispatch({
        //     type: 'graph/queryGraphs'
        // });
        //console.log(mapStateToProps(props).graph)
    }}>aaa

  </div>
};


const mapStateToProps = (state) =>{
    return {
        graph:state.graph,
    }
}

export default connect(mapStateToProps)(HomePage);
