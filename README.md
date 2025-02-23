# dpattern
design pattern

## options pattern(选项模式)
### 优点
1. 支持传递多个参数，并且在参数发生变化时保持兼容性
2. 支持任意顺序传递参数
3. 支持默认值
4. 方便扩展
5. 通过 WithXXX 的函数命名，可以使参数意义更加明确

### 缺点
为了实现选项模式，增加了很多代码

### 适合场景
1. 结构体参数很多，创建结构体时，期望创建一个携带默认值的结构体变量，并选择性修改其中一些参数的值。
2. 构体参数经常变动，变动时又不想修改创建实例的函数。例如：结构体新增一个 retry 参数，但是又不想在 NewConnect 入参列表中添加 retry int 这样的参数声明。

## template pattern(模版模式)
### 介绍
模板模式 (Template Pattern) 定义一个操作中算法的骨架，而将一些步骤延迟到子类中。这种方法让子类在不改变一个算法结构的情况下，就能重新定义该算法的某些特定步骤。

简单来说，模板模式就是将一个类中能够公共使用的方法放置在抽象类中实现，将不能公共使用的方法作为抽象方法，强制子类去实现，这样就做到了将一个类作为一个模板，让开发者去填充需要填充的地方。

### 优点
1. 封装不变部分，扩展可变部分。模板方法模式将可变的部分封装在抽象方法中，不变的部分封装在基本方法中。这使得子类可以根据需求对可变部分进行扩展，而不变部分仍然保持不变。
2. 避免重复代码，抽象类中包含的基本方法可以避免子类重复实现相同的代码逻辑。
3. 更好的扩展性，由于具体实现由子类来完成，因此可以方便地扩展新的功能或变更实现方式，同时不影响模板方法本身。

### 缺点
1. 类多，由于每个算法都需要一个抽象类和具体子类来实现，因此在操作流程比较多时可能导致类的数量急剧增加，从而导致代码的复杂性提高。
2. 关联性高，模板方法与子类实现的抽象方法紧密相关，如果该模板方法需要修改，可能会涉及到多个子类的修改。

### 应用场景
1. 开发框架，通常框架会定义一些通用的模板，子类可以根据自身的特定需求来细化模板的实现细节。
2. 业务逻辑，可以针对业务流程做一些拆解，将特定步骤改为子类实现。比如发送验证码的流程，在发送验证码时需要选择不同厂商来发送验证码，但是发送的验证码前的检查、验证码生成、保存验证码逻辑都是一样的。

## factory pattern(工厂模式)
### 简单工厂
简单工厂模式是最常用、最简单的。它就是一个接受一些参数，然后返回结构体实例的函数

### 抽象工厂
和简单工厂模式的唯一区别，就是它返回的是接口而不是结构体。

### 工厂方法
在**简单工厂模式**中，依赖于唯一的工厂对象，如果我们需要实例化一个产品，就要向工厂中传入一个参数，获取对应的对象；如果要增加一种产品，就要在工厂中修改创建产品的函数。
这会导致耦合性过高，这时我们就可以使用**工厂方法**模式。

在**工厂方法模式**中，依赖工厂函数，我们可以通过实现工厂函数来创建多种工厂，将对象创建从由一个对象负责所有具体类的实例化，变成由一群子类来负责对具体类的实例化，从而将过程解耦。

## strategy pattern(策略模式)
策略模式（Strategy Pattern）定义一组算法，将每个算法都封装起来，并且使它们之间可以互换。

### 应用场景
> 经常要根据不同的场景，采取不同的措施，也就是不同的策略

比如，假设我们需要对 a、b 这两个整数进行计算，根据条件的不同，需要执行不同的计算方式。我们可以把所有的操作都封装在同一个函数中，然后通过 if ... else ... 的形式来调用不同的计算方式，这种方式称之为硬编码。

在实际应用中，随着功能和体验的不断增长，我们需要经常添加 / 修改策略，这样就需要不断修改已有代码，不仅会让这个函数越来越难维护，还可能因为修改带来一些 bug。
所以为了解耦，需要使用策略模式，定义一些独立的类来封装不同的算法，每一个类封装一个具体的算法（即策略）。

## proxy pattern(代理模式)
代理模式 (Proxy Pattern)，可以为另一个对象提供一个替身或者占位符，以控制对这个对象的访问。

> 结构
1. 抽象主题（Subject）:定义了真实主题和代理主题的共同接口，这样在任何使用真实主题的地方都可以使用代理主题。
2. 真实主题（Real Subject）: 实现了抽象主题接口，是代理对象所代表的真实对象。客户端直接访问真实主题，但在某些情况下，可以通过代理主题来间接访问。
3. 代理（Proxy）: 实现了抽象主题接口，并持有对真实主题的引用。代理主题通常在真实主题的基础上提供一些额外的功能，例如延迟加载、权限控制、日志记录等。
4. 客户端（Client）: 使用抽象主题接口来操作真实主题或代理主题，不需要知道具体是哪一个实现类。

> 主要解决的问题
1. 远程代理（Remote Proxy）：在分布式系统中，代理模式可以用来实现远程代理，允许一个对象在不同的地址空间中进行通信。比如，你可以使用代理对象来代表一个位于不同服务器上的对象，通过网络进行通信，Go 语言中常用 gRPC 或者 HTTP 实现这种远程代理。
2. 虚拟代理（Virtual Proxy）：虚拟代理是延迟对象的创建，直到真正需要时。这在 Go 中可以用于加载大对象或者延迟初始化。例如，一个图片浏览器可以使用虚拟代理来延迟加载高分辨率图片，而不是在初始化时就加载所有图片。
3. 保护代理（Protection Proxy）：保护代理控制对原始对象的访问。这在权限控制和安全性方面特别有用。例如，一个保护代理可以验证调用者是否具有执行操作的权限，然后再决定是否将请求转发给真实的对象。
4. 缓存代理（Caching Proxy）：缓存代理用于为开销大的操作结果提供缓存，以便多个客户端可以共享结果。在 Go 中，可以使用代理模式来实现缓存代理，以加快数据访问速度并减少对底层资源的频繁访问。

> 优点
1. 职责分离：代理模式将访问控制与业务逻辑分离。
2. 扩展性：可以灵活地添加额外的功能或控制。
3. 智能化：可以智能地处理访问请求，如延迟加载、缓存等。

> 缺点
1. 性能开销：增加了代理层可能会影响请求的处理速度。
2. 实现复杂性：某些类型的代理模式实现起来可能较为复杂。

> 注意事项
1. 与适配器模式的区别：适配器模式改变接口，而代理模式不改变接口。
2. 与装饰器模式的区别：装饰器模式用于增强功能，代理模式用于控制访问。

## singleton pattern(单例模式)
单例模式（Singleton Pattern），是最简单的一个模式。在 Go 中，单例模式指的是全局只有一个实例，并且它负责创建自己的对象。单例模式不仅有利于减少内存开支，还有减少系统性能开销、防止多个实例产生冲突等优点。

因为单例模式保证了实例的全局唯一性，而且只被初始化一次，所以比较适合全局共享一个实例，且只需要被初始化一次的场景，例如数据库实例、全局配置、全局任务池等。

单例模式又分为饿汉方式和懒汉方式。饿汉方式指全局的单例实例在包被加载时创建，而懒汉方式指全局的单例实例在第一次被使用时创建。你可以看到，这种命名方式非常形象地体现了它们不同的特点。

## decorator pattern(装饰器模式)
装饰器模式（Decorator Pattern）允许向一个现有的对象添加新的功能，同时又不改变其结构。这种类型的设计模式属于结构型模式，它是作为现有的类的一个包装。

装饰器模式通过将对象包装在装饰器类中，以便动态地修改其行为。

这种模式创建了一个装饰类，用来包装原有的类，并在保持类方法签名完整性的前提下，提供了额外的功能。

> 解决问题

1. 避免通过继承引入静态特征，特别是在子类数量急剧膨胀的情况下。
2. 允许在运行时动态地添加或修改对象的功能。

> 使用场景

1. 当需要在不增加大量子类的情况下扩展类的功能。

2. 当需要动态地添加或撤销对象的功能。

> 优点

1. 低耦合：装饰类和被装饰类可以独立变化，互不影响。
2. 灵活性：可以动态地添加或撤销功能。
3. 替代继承：提供了一种继承之外的扩展对象功能的方式。

> 缺点

复杂性：多层装饰可能导致系统复杂性增加。

> 使用建议

1. 在需要动态扩展功能时，考虑使用装饰器模式。
2. 保持装饰者和具体组件的接口一致，以确保灵活性。

## 相关链接
菜鸟教程：https://www.runoob.com/design-pattern/design-pattern-tutorial.html