package converter

import (
	"fmt"
	"testing"
)

func TestHtmlConverter_Convert(t *testing.T) {
	var markdownStr = `<p><strong>参考</strong>：
<a href="https://blog.csdn.net/sun_promise/article/details/51315032">https://blog.csdn.net/sun_promise/article/details/51315032</a><br />
<a href="https://docs.oracle.com/javase/tutorial/java/annotations/type_annotations.html">https://docs.oracle.com/javase/tutorial/java/annotations/type_annotations.html</a><br />
<a href="https://docs.oracle.com/javase/tutorial/java/annotations/repeating.html">https://docs.oracle.com/javase/tutorial/java/annotations/repeating.html</a>  </p>
<h2>概念</h2>
<p>Java提供了一种原程序中的元素关联任何信息和任何元数据的途径和方法。（注解就是元数据）</p>
<blockquote>
<p>JDK1.5 提供注解特性。</p>
</blockquote>
<h2>元注解</h2>
<p>修饰注解的注解，称为元注解，在JDK中，共有四种元注解：</p>
<table>
<thead>
<tr>
<th>名称</th>
<th>作用</th>
</tr>
</thead>
<tbody>
<tr>
<td><code>@Target</code></td>
<td>规定注解可以使用的位置，类，方法还是域..</td>
</tr>
<tr>
<td><code>@Retention</code></td>
<td>表示需要在什么级别保存该注解信息，由RetentionPolicy枚举定义</td>
</tr>
<tr>
<td><code>@Documented</code></td>
<td>表示注解会被包含在javaapi文档中</td>
</tr>
<tr>
<td><code>@Inherited</code></td>
<td>允许子类继承父类的注解</td>
</tr>
</tbody>
</table>
<p><strong>@Target</strong>：</p>
<blockquote>
<p>使用ElementType枚举定义，包含以下值：</p>
<ul>
<li>CONSTRUCTOR：构造器的声明</li>
<li>FIELD：域声明（包括enum实例）</li>
<li>LOCAL_VARIABLE：局部变量声明</li>
<li>METHOD：方法声明</li>
<li>PACKAGE：包声明</li>
<li>PARAMETER：参数声明</li>
<li>TYPE：类、接口（包括注解类型）或enum声明</li>
<li>ANNOTATION_TYPE：注解声明（应用于另一个注解上）</li>
<li>TYPE_PARAMETER：类型参数声明（1.8新加入）</li>
<li>TYPE_USE：类型使用声明（1.8新加入）</li>
<li>PS:当注解未指定Target值时，此注解可以使用任何元素之上，就是上面的类型</li>
</ul>
</blockquote>
<p><strong>@Retention</strong>：</p>
<blockquote>
<p>使用RetentionPolicy枚举定义，包含以下值：</p>
<ul>
<li>SOURCE：注解将被编译器丢弃（该类型的注解信息只会保留在源码里，源码经过编译后，注解信息会被丢弃，不会保留在编译好的class文件里）</li>
<li>CLASS：注解在class文件中可用，但会被VM丢弃（该类型的注解信息会保留在源码里和class文件里，在执行的时候，不会加载到虚拟机（JVM）中）</li>
<li>RUNTIME：VM将在运行期也保留注解信息，因此可以通过反射机制读取注解的信息（源码、class文件和执行的时候都有注解的信息）
PS：当注解未定义Retention值时，默认值是CLASS</li>
</ul>
</blockquote>
<h2>常用注解</h2>
<ul>
<li><code>java.lang.SuppressWarnings</code> 抑制警告</li>
<li><code>java.lang.Override</code> 方法重写</li>
<li><code>java.lang.Deprecated</code> 已经过时</li>
</ul>
<p>jdk7新增：</p>
<ul>
<li><code>@SafeVarargs</code> 堆污染</li>
</ul>
<h2>定义注解</h2>
<pre><code class="language-java">// 定义的注解默认会继承java.lang.annotation.Annotation接口
public @interface 注解名{

}</code></pre>
<h2>底层实现</h2>
<p>注解的本质实际上是一个Interface，所有注解在反编译后都是继承<code>java.lang.annotation.Annotation</code>接口的。</p>
<h2>类型注解(Type Annotations)</h2>
<p>在JDK8之前，类型注解只能用在类型声明的地方：</p>
<pre><code class="language-java">@MyAnnotation // 声明类
public class Test {
}</code></pre>
<p>JDK8后，类型注解可以在任何有类型的地方加入注解(包含泛型)：</p>
<pre><code class="language-java">@MyAnnotation
public class Test&lt;@MyAnnotation T&gt; extends @MyAnnotation Fathers implements @MyAnnotation Callable {

    private @MyAnnotation String name = &quot;lisi&quot;;

    public @MyAnnotation String test(@MyAnnotation String name) throws @MyAnnotation Exception {
        @MyAnnotation String s = &quot;test&quot;;
        @MyAnnotation List&lt;String&gt; list = new @MyAnnotation ArrayList&lt;&gt;();
        Number a = 1;
        Integer i = (@MyAnnotation Integer) a;
        if (list.get(0) == null) {
            throw new @MyAnnotation Exception();
        }
        list.add(s);
        return list.get(0);
    }

    @Override
    public Object call() throws Exception {
        return null;
    }
}</code></pre>
<h3>作用</h3>
<p>在8种新增类型注释主要是为了改进Java的程序分析，配合类型检查框架做强类型检查，从而在编译期间确认运行时异常，比如 NullpointException，从而提高代码质量。
比如：</p>
<pre><code class="language-java">@NonNull Object my = null;</code></pre>
<p>第三方工具会在编译期间自动检测my是否为null，如果为null，抛出异常或者警告</p>
<blockquote>
<p>目前具有此项功能的检测框架有 <a href="https://checkerframework.org/">The Check Framework</a></p>
</blockquote>
<h3>创建类型注解</h3>
<p>JDK8 新增<code>ElementType.TYPE_USE</code>用来创建类型注解：</p>
<pre><code class="language-java">@Target({TYPE_USE})
@Retention(RUNTIME)
public @interface MyAnnotation {
}</code></pre>
<p><strong>注意</strong>：JDK8还提供了 TYPE_PARAMETER 类型的注解，表示可以修饰类型参数（泛型）的注解。</p>
<h2>可重复注解(Repeating Annotations)</h2>
<p>在8之前，一个目标只可以打一个注解：</p>
<pre><code class="language-java">@Test
private String name;</code></pre>
<p>现在JDK8提供了Repeating Annotations，可以在一个目标上打上多个可重复注解：</p>
<pre><code class="language-java">@Test(&quot;a&quot;)
@Test(&quot;b&quot;)
private String name;</code></pre>
<h3>创建可重复注解</h3>
<p>JDK8中提供了一个新的源注解，用于表示可重复注解：</p>
<pre><code class="language-java">// 创建可重复注解：指定容器类型为 Sechedules.class
@Repeatable(Schedules.class)
@Retention(RetentionPolicy.RUNTIME)
public @interface Schedule {
    String dayOfMonth() default &quot;first&quot;;
    String dayOfWeek() default &quot;Mon&quot;;
    int hour() default 12;
}
// 可重复注解的容器注解，用于存放多个可重复注解： 
@Retention(RetentionPolicy.RUNTIME)
public @interface Schedules {
    Schedule[] value();
}</code></pre>
<h3>获取可重复注解</h3>
<p>注意：如果想要在运行时获取注解，注解<code>Retention</code>必须为<code>RetentionPolicy.RUNTIME</code>，并且，Schedules与Schedule 注解的保留时期必须都为运行时！</p>
<p>获取可重复注解的方式同获取普通注解没有什么太大的区别：</p>
<pre><code class="language-java">@Schedule(hour = 9)
@Schedule(hour = 10)
public class Task {

    public static void main(String[] args) throws Exception {
        Schedule[] ats = Task.class.getAnnotationsByType(Schedule.class);
        System.out.println(Arrays.toString(ats));  
        // [@test.s.Schedule(hour=9, dayOfMonth=first, dayOfWeek=Mon), @test.s.Schedule(hour=10, dayOfMonth=first, dayOfWeek=Mon)]

        System.out.println(Arrays.toString(Task.class.getAnnotations())); 
        // [@test.s.Schedules(value=[@test.s.Schedule(hour=9, dayOfMonth=first, dayOfWeek=Mon), @test.s.Schedule(hour=10, dayOfMonth=first, dayOfWeek=Mon)])]
        //注意：这里获取到的是 Schedules 注解，而不是多个 Schedule 注解
    }
}</code></pre>
<h3>可重复注解的原理</h3>
<p>使用一个注解来存储重复的注解，编译后的class为：</p>
<pre><code class="language-java">@Schedules({@Schedule(hour = 9), @Schedule(hour = 10)})
public class Task {
    public Task() {
    }
    public static void main(String[] args) throws Exception {
        Schedule[] ats = (Schedule[])Task.class.getAnnotationsByType(Schedule.class);
        System.out.println(Arrays.toString(ats));
        System.out.println(Arrays.toString(Task.class.getAnnotations()));
    }
}</code></pre>
<p>所谓的可重复注解，<strong>只是编译层面的改动</strong>。</p>`
	c := NewConverter(HTML)
	r, _ := c.Convert([]byte(markdownStr))
	fmt.Println("----")
	fmt.Println(string(r))

}
