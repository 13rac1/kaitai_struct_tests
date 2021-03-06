using NUnit.Framework;

namespace Kaitai
{
    [TestFixture]
    public class SpecParamsCallShort : CommonSpec
    {
        [Test]
        public void TestParamsCallShort()
        {
            var r = ParamsCallShort.FromFile(SourceFile("term_strz.bin"));
            Assert.AreEqual(r.Buf1.Body, "foo|b");
            Assert.AreEqual(r.Buf2.Body, "ar|ba");
            Assert.AreEqual(r.Buf2.Trailer, 0x7a);
        }
    }
}
