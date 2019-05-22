using Org.BouncyCastle.Crypto;
using Org.BouncyCastle.Security;
using System;
using System.Security.Cryptography;
using System.Text;

namespace ConsoleApp1
{
    public class BigerSign
    {
        public string RsaWithSha256Sign(string queryString, string method, string expiry, string body)
        {
            var data = $"{queryString}{method}{expiry}{body}";

            IBufferedCipher c = CipherUtilities.GetCipher("RSA/ECB/PKCS1Padding");// 参数与Java中加密解密的参数一致

            //第一个参数为true表示加密，为false表示解密；第二个参数表示密钥  
            c.Init(true, GetPrivateKeyParameter("myprivateKey"));
            byte[] DataToEncrypt = Encoding.UTF8.GetBytes(data);
            byte[] hash = SHA256Managed.Create().ComputeHash(DataToEncrypt);
            byte[] outBytes = c.DoFinal(hash);//加密  
            string strBase64 = Convert.ToBase64String(outBytes);
            return strBase64;
        }

        private AsymmetricKeyParameter GetPrivateKeyParameter(string privateKey)
        {
            string prikey = privateKey.Replace("\r", "").Replace("\n", "").Replace(" ", "");
            byte[] privateInfoByte = Convert.FromBase64String(prikey);
            AsymmetricKeyParameter priKey = PrivateKeyFactory.CreateKey(privateInfoByte);
            return priKey;
        }
    }
}
