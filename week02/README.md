## 第二周作业
Q：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

A：dao层是业务层的最底层，用来直接和数据库进行交互，dao层遇到 error 时建议 wrap error，最终如何处理由业务层进行处理，然后根据不同的应用场景进行处理
