# panic

GoでのExceptionがpanic  
panicはrecover()でcatchする事が可能です。  
どこで起きるか分からないので、deferで必ずキャッチできる様にする必要があります。  
recoverしない場合は、プロセスが終了します。  
コードではpanic例をいくつか記載しました。  
panicは基本的に扱わず、自分で`panic()`を呼ぶ事は殆どありません。  
