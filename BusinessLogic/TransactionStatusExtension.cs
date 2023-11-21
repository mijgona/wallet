using DataAccess;

namespace BusinessLogic;
 
public static class TransactionStatusExtension
{
    public static TransactionStatus ToTransactionStatusEnum(this string typeStr)
        => Enum.Parse<TransactionStatus>(typeStr);
}