# GoRedistributableChecker

### Supported redistributables
- Microsoft Visual C++ Redistributable Package 2015 - 2019 (x86) :heavy_check_mark:
- Microsoft Visual C++ Redistributable Package 2015 - 2019 (x64) :heavy_check_mark:
- Microsoft Visual C++ Redistributable Package 2017 (x86) :heavy_check_mark:
- Microsoft Visual C++ Redistributable Package 2017 (x64) :heavy_check_mark:
- Microsoft Visual C++ Redistributable Package 2015 (x86) :heavy_check_mark:
- Microsoft Visual C++ Redistributable Package 2015 (x64) :heavy_check_mark:
- Microsoft Visual C++ Redistributable Package 2013 (x86) :heavy_check_mark:
- Microsoft Visual C++ Redistributable Package 2013 (x64) :heavy_check_mark:
- Microsoft Visual C++ Redistributable Package 2012 (x86) :heavy_check_mark:
- Microsoft Visual C++ Redistributable Package 2012 (x64) :heavy_check_mark:
- Microsoft Visual C++ Redistributable Package 2010 (x86) :heavy_check_mark:
- Microsoft Visual C++ Redistributable Package 2010 (x64) :heavy_check_mark:
- Microsoft Visual C++ Redistributable Package 2008 (x86) :heavy_check_mark:
- Microsoft Visual C++ Redistributable Package 2008 (x64) :heavy_check_mark:
- Microsoft Visual C++ Redistributable Package 2005 (x86) :heavy_check_mark:
- Microsoft Visual C++ Redistributable Package 2005 (x64) :heavy_check_mark:

### Code Sample
``` 
import redistributable;

if redistributable.IsInstalled(redistributable.VC2017x64) {
  //go on
}
```

### Notes
Worked on my **Windows 10 (x64)**

Derived from : https://github.com/bitbeans/RedistributableChecker

Windows Registry Keys taken from: https://stackoverflow.com/a/34209692/4013391
