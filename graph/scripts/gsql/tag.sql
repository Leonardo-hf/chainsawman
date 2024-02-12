INSERT INTO graph.algo(id, name, define, detail, isTag) VALUES (1, '其他', '一系列从不同角度研究网络性质的算法。', null, 1);

INSERT INTO graph.algo(id, name, define, detail, isTag) VALUES (2, '软件研发影响力', '软件研发影响力是对一个软件对其他软件功能的影响大小的度量。', '软件研发影响力是对一个软件对其他软件功能的影响大小的度量。具体表现为开发者在实现软件功能的过程中是否须要依赖某些第三方软件以及更倾向于使用哪些第三方软件。具备高研发影响力的软件应当表现为被开发者广泛地使用或间接使用以开发其软件功能。一个具有高研发影响力的开源软件发生闭源或被废弃，将对大量其他软件的功能的可靠性造成负面影响。

进一步地，我们探讨软件在开发过程中如何被开发者使用。若开发者为了实现一个软件功能需要使用数据库，而访问数据库时则必须使用数据库驱动，数据库驱动对该软件的研发具有影响力，而开发者在开发过程中使用Lint工具来保障软件质量，这个工具对软件功能本身没有帮助，则可以认为Lint工具对软件的研发不具有影响力。我们认为具备研发影响力的软件可能存在四种性质，

1. 基础性质的软件，例如工具库等在功能开发中被普遍使用的软件
2. 底层性质的软件，例如网络协议等往往在被间接使用的软件
3. 框架性质的软件，例如开发框架等集成了众多功能的软件
4. 生态性质的软件，即为了使用软件的相关生态而须要使用的软件

假设一：具有基础、底层、框架、生态性质的软件相比其他软件更具备研发影响力

进一步地，我们基于依赖图谱，从软件被依赖的拓扑形式研究软件在开发中扮演的角色

假设二：入度越大的软件越具备基础性质

假设三：在更多的路径上处于尾部的软件具有更多的底层性质

假设四：在更多的路径上处于中间的软件具有更多的框架性质

假设五：对依赖网络稳定性贡献越大的软件具有更多的生态性质

进一步，基于专家经验AHP方法，我们认为：

假设五：具有基础性质的软件的研发影响力 > 具有底层性质的软件的研发影响力 > 具有生态性质的软件的研发影响力 > 具有框架性质的软件的研发影响力

进一步，对指标识别的高影响力软件，进一步分类，发现每种指标识别出的软件既在主观上对研发很有帮助（很常见），也符合开发者的认知（第三方评估指标），此外，还符合假设中提及的软件在开发中扮演的角色（比如基础类库属于基础性质，网络协议属于底层性质）。', 1);

INSERT INTO graph.algo(id, name, define, detail, isTag) VALUES (3, '软件卡脖子风险', '卡脖子软件指发生断供后会对下游依赖组织造成重大影响且该影响无法在短时间内消除的软件，软件卡脖子风险算法即计算开源软件成为行业内卡脖子软件的可能性。', '卡脖子软件指发生断供后会对下游依赖组织造成重大影响且该影响无法在短时间内消除的软件。为满足以上特征，我们认为一个卡脖子软件需要具备以下三种特征。

* 假设1：卡脖子软件需要具有高商业价值，通过使用该软件，组织可以提高自身在市场上的竞争力，具有高商业价值的软件断供会对竞争对手造成有力打击。
  * 假设1-1：代表着尖端技术、具有高产值的软件具有高商业价值。

  * 假设1-2：生产力工具、能够帮助提高生产效率的软件具有高商业价值。

* 假设2：卡脖子软件具有高开发成本，这会导致造成的危害得以持续。
  * 假设2-1：代码量大、组件结构复杂的软件具有高开发成本。

  * 假设2-2：有系列配套软件及活跃下游用户的软件具有高开发成本。

* 假设3：卡脖子软件需要具备实施断供的条件。

  * 假设3-1：符合供给关系的软件才可以卡脖子，即只有供应方对被供应方卡脖子。

  * 假设3-2：软件团队内部分人员或第三方组织有强话语权才能实施卡脖子。

![](/assets/strangle-procedure.png)', 1);

