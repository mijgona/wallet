using DataAccess;

namespace BusinessLogic;
 
    public static class TransactionTypeExtension
    {
        public static TransactionType ToTransactionTypeEnum(this string typeStr)
            => Enum.Parse<TransactionType>(typeStr);
    }