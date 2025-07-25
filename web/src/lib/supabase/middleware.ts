import type { NextRequest } from "next/server"
import { createServerClient } from "@supabase/ssr"
import { NextResponse } from "next/server"

export async function updateSession(request: NextRequest, response: NextResponse) {
  const supabase = createServerClient(
    process.env.NEXT_PUBLIC_SUPABASE_URL!,
    process.env.NEXT_PUBLIC_SUPABASE_ANON_KEY!,
    {
      cookies: {
        getAll() {
          const cookies = request.cookies.getAll()
          return cookies
        },
        setAll(cookiesToSet) {
          cookiesToSet.forEach(({ name, value }) => request.cookies.set(name, value))
          const res = NextResponse.next({ request })
          cookiesToSet.forEach(({ name, value, options }) =>
            res.cookies.set(name, value, options),
          )
        },
      },
    },
  )

  try {
    const { data: { user } } = await supabase.auth.getUser()
    return { response, user }
  } catch {
    return { response, user: null }
  }
}
