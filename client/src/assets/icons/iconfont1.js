/* eslint-disable */
!(function (d) {
    var t,
        n =
            '<svg><symbol id="icon-liebiao" viewBox="0 0 1024 1024"><path d="M892.928 128q28.672 0 48.64 19.968t19.968 48.64l0 52.224q0 28.672-19.968 48.64t-48.64 19.968l-759.808 0q-28.672 0-48.64-19.968t-19.968-48.64l0-52.224q0-28.672 19.968-48.64t48.64-19.968l759.808 0zM892.928 448.512q28.672 0 48.64 19.968t19.968 48.64l0 52.224q0 28.672-19.968 48.64t-48.64 19.968l-759.808 0q-28.672 0-48.64-19.968t-19.968-48.64l0-52.224q0-28.672 19.968-48.64t48.64-19.968l759.808 0zM892.928 769.024q28.672 0 48.64 19.968t19.968 48.64l0 52.224q0 28.672-19.968 48.64t-48.64 19.968l-759.808 0q-28.672 0-48.64-19.968t-19.968-48.64l0-52.224q0-28.672 19.968-48.64t48.64-19.968l759.808 0z"  ></path></symbol><symbol id="icon-xiazai" viewBox="0 0 1024 1024"><path d="M819.203 405.649c0-169.66-137.541-307.19-307.201-307.19s-307.195 137.53-307.195 307.19c-113.105 0-204.8 91.69-204.8 204.801s91.695 204.801 204.8 204.801h102.4V733.33h-102.4c-67.755 0-122.88-55.12-122.88-122.88 0-67.761 55.125-122.881 122.88-122.881h81.92v-81.92c0-124.22 101.055-225.28 225.275-225.28 124.221 0 225.281 101.06 225.281 225.28v81.92h81.92c67.76 0 122.871 55.12 122.871 122.881 0 67.76-55.111 122.88-122.871 122.88h-102.4v81.921h102.4c113.09 0 204.791-91.69 204.791-204.801s-91.701-204.801-204.791-204.801z" fill="#040000" ></path><path d="M511.393 925.541l221.281-238.02-64.441-60-110.79 119.22V410.22h-92.16v336.47L354.488 627.521l-64.431 60z" fill="#040000" ></path></symbol></svg>',
        e = (t = document.getElementsByTagName("script"))[
            t.length - 1
        ].getAttribute("data-injectcss");
    if (e && !d.__iconfont__svg__cssinject__) {
        d.__iconfont__svg__cssinject__ = !0;
        try {
            document.write(
                "<style>.svgfont {display: inline-block;width: 1em;height: 1em;fill: currentColor;vertical-align: -0.1em;font-size:16px;}</style>"
            );
        } catch (t) {
            console && console.log(t);
        }
    }
    !(function (t) {
        if (document.addEventListener)
            if (
                ~["complete", "loaded", "interactive"].indexOf(
                    document.readyState
                )
            )
                setTimeout(t, 0);
            else {
                var e = function () {
                    document.removeEventListener("DOMContentLoaded", e, !1),
                        t();
                };
                document.addEventListener("DOMContentLoaded", e, !1);
            }
        else
            document.attachEvent &&
                ((n = t),
                (i = d.document),
                (o = !1),
                (l = function () {
                    o || ((o = !0), n());
                }),
                (c = function () {
                    try {
                        i.documentElement.doScroll("left");
                    } catch (t) {
                        return void setTimeout(c, 50);
                    }
                    l();
                })(),
                (i.onreadystatechange = function () {
                    "complete" == i.readyState &&
                        ((i.onreadystatechange = null), l());
                }));
        var n, i, o, l, c;
    })(function () {
        var t, e;
        ((t = document.createElement("div")).innerHTML = n),
            (n = null),
            (e = t.getElementsByTagName("svg")[0]) &&
                (e.setAttribute("aria-hidden", "true"),
                (e.style.position = "absolute"),
                (e.style.width = 0),
                (e.style.height = 0),
                (e.style.overflow = "hidden"),
                (function (t, e) {
                    e.firstChild
                        ? (function (t, e) {
                              e.parentNode.insertBefore(t, e);
                          })(t, e.firstChild)
                        : e.appendChild(t);
                })(e, document.body));
    });
})(window);
